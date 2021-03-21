package test

import (
	"context"
	"fmt"
	_"io/fs"
	_"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"
	_"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/draftms/go_library/logging"
)

var logger = logging.NewInstance()
var clientOptions *options.ClientOptions
var client *mongo.Client
var ctx context.Context

type Analyzer struct {
	ID primitive.ObjectID 	`bson:"_id,omitempty"`
	Name string 			`bson:"name,omitempty"`
	Code string 			`bson:"code,omitempty"`
	Cusvalue string 		`bson:"cusvalue,omitempty"`
	Tags []string			`bson:"tags,omitempty"`
}

type ExpectedStudy struct {
	ID primitive.ObjectID	`bson:"_id,omitempty"`
	RegistDate string		`bson:"rD,omitempty"`
	StudyDate string		`bson:"sD,omitempty"`
	HospitalName string		`bson:"h,omitempty"`
	RequestYN string		`bson:"rY,omitempty"`
	BackupYN string			`bson:"bY,omitempty"`
	StudyUID string			`bson:"stU,omitempty"`
	SeriesUID string		`bson:"seC,omitempty"`
	AddtionalBackupF string	`bson:"abf,omitempty"`
	MRequestY string 		`bson:"mrY,omitempty"`
	UploadInfo []UploadInfoData 	`bson:"uploadInfo,omitempty"`
}

type UploadInfoData struct {
	S3Key string			`bson:"s3key,omitempty"`
	Size int				`bson:"size,omitempty"`
	BackupDate string		`bson:"bD,omitempty"`
	SOPCount int			`bson:"soC,omitempty"`
}

func TestExpectedStudy(t *testing.T) {
	//Find mongoDB data
	mongoDBConnect()
    collection := client.Database("Dev").Collection("ExpectedStudy")

	findResultCursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		logger.Fatal(err)
	}

	if err = findResultCursor.All(ctx, bson.D{}); err != nil {
		fmt.Println(err)
	}

	file, err := os.Create("modifiedUploadInfoFile.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for findResultCursor.Next(ctx) {
		var findResult ExpectedStudy
		err := findResultCursor.Decode(&findResult)

		if err != nil {
			fmt.Println("cursor.Next() error :",err)
			os.Exit(1)
		} else {
			fmt.Println("result type :", reflect.TypeOf(findResult))
			fmt.Println("result :", findResult)

			haveDCMFile := false

			var prvUploadInfo UploadInfoData
			var newUploadInfo UploadInfoData

			newExpectedStudy := findResult
			newExpectedStudy.UploadInfo = nil

			for idx, uploadInfoValue := range findResult.UploadInfo {
				if strings.Contains(uploadInfoValue.S3Key,".dcm") {
					haveDCMFile = true
				}

				if idx == 0 {
					prvUploadInfo = uploadInfoValue
				} else if !haveDCMFile && idx == 1 {
					if uploadInfoValue.S3Key == prvUploadInfo.S3Key && 
				   	   uploadInfoValue.Size == prvUploadInfo.Size &&
					   uploadInfoValue.BackupDate == prvUploadInfo.BackupDate {

						newUploadInfo.S3Key = uploadInfoValue.S3Key
						newUploadInfo.Size = uploadInfoValue.Size
						newUploadInfo.BackupDate = uploadInfoValue.BackupDate
						newUploadInfo.SOPCount = uploadInfoValue.SOPCount + prvUploadInfo.SOPCount

						newExpectedStudy.UploadInfo = append(newExpectedStudy.UploadInfo, newUploadInfo)
					   }
				}
			}

			if !haveDCMFile {
				//uploadInfo 를 하나로 병합
				updateResult, err := collection.UpdateOne(
					ctx, 
					bson.M{"_id": findResult.ID},
					bson.D{
						{"$set", bson.D{{"uploadInfo",newExpectedStudy.UploadInfo[0]}}},
					},
				)
	
				if err != nil {
					fmt.Println("UpdateErr")
				}
	
				fmt.Printf("Update %v Document\n", updateResult.ModifiedCount)

				//업데이트 doc를 파일로 저장
				fmt.Fprintln(file, findResult.ID.Hex())
			}
		}
	}

	client.Disconnect(ctx)
}

func mongoDBConnect(){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

    clientOptions = options.Client().ApplyURI("mongodb://localhost:27017")

	var err error
    client, err = mongo.Connect(ctx, clientOptions)

    if err != nil { 
        logger.Error(err)
    }

    err = client.Ping(ctx, nil)

    if err != nil {
        logger.Error(err)
    }
}

func TestMongoDBInsertOne(t *testing.T) {
	mongoDBConnect()
    collection := client.Database("Dev").Collection("analyzer")

	//1.Insert mongoDB data
 	analyzer_doc := Analyzer{
		Name:"test2", 
		Code:"10-16", 
		Cusvalue: "hud2",
		Tags: []string{"a","b","c"},
	} 
	insertResult, err := collection.InsertOne(ctx, analyzer_doc)

	if err != nil {
		logger.Fatal(err)
	}

    fmt.Println("Add data : ", insertResult.InsertedID)

	client.Disconnect(ctx)
}

func TestMongoDBInsertMany(t *testing.T) {
	mongoDBConnect()
    collection := client.Database("Dev").Collection("analyzer")

	//1.Insert mongoDB data
 	analyzer_doc1 := Analyzer{
		Name:"test2", 
		Code:"10-16", 
		Cusvalue: "hud2",
		Tags: []string{"a","b","c"},
	} 

	analyzer_doc2 := Analyzer{
		Name:"test2", 
		Code:"10-16", 
		Cusvalue: "hud2",
		Tags: []string{"a","b","c"},
	}

	analyzeDataList := []interface{}{analyzer_doc1, analyzer_doc2}
	insertResult, err := collection.InsertMany(ctx, analyzeDataList)

	if err != nil {
		logger.Fatal(err)
	}

    fmt.Println("Add data : ", insertResult.InsertedIDs)

	client.Disconnect(ctx)
}

func TestMongoDBFindAll(t *testing.T) {
	//Find mongoDB data
	mongoDBConnect()
    collection := client.Database("Dev").Collection("analyzer")

	findResultCursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		logger.Fatal(err)
	}

	if err = findResultCursor.All(ctx, bson.D{}); err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Add data : ", findResult.ID())

	for findResultCursor.Next(ctx) {
		var result bson.M
		err := findResultCursor.Decode(&result)

		if err != nil {
			fmt.Println("cursor.Next() error :",err)
			os.Exit(1)
		} else {
			fmt.Println("result type :", reflect.TypeOf(result))
			fmt.Println("result :", result)

			var analyzerData Analyzer
			bsonBytes, _ := bson.Marshal(result)
			bson.Unmarshal(bsonBytes, &analyzerData)

			fmt.Println("struct data :", result)
		}
	}

	client.Disconnect(ctx)
}

func TestMongoDBFindOne(t *testing.T) {
	//Find mongoDB data
	mongoDBConnect()
    collection := client.Database("Dev").Collection("analyzer")

	findResultCursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		logger.Fatal(err)
	}

	if err = findResultCursor.All(ctx, bson.D{}); err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Add data : ", findResult.ID())

	for findResultCursor.Next(ctx) {
		var result bson.M
		err := findResultCursor.Decode(&result)

		if err != nil {
			fmt.Println("cursor.Next() error :",err)
			os.Exit(1)
		} else {
			fmt.Println("result type :", reflect.TypeOf(result))
			fmt.Println("result :", result)
		}
	}

	client.Disconnect(ctx)
}

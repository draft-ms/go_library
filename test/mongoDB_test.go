package test

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

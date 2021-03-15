package main_test

import (
	"testing"
	"context"
	"fmt"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/draftms/go_library/logging"
)

var logger = logging.NewInstance()
var clientOptions *options.ClientOptions
var client *mongo.Client

func mongoDBConnect(){

    clientOptions = options.Client().ApplyURI("mongodb://localhost:27017/admin?safe=true")

	var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)

    if err != nil { 
        logger.Error(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        logger.Error(err)
    }

}

func TestMongoDBInsert(t *testing.T) {
	mongoDBConnect()
    collection := client.Database("local").Collection("analyzer")

    //var result user

    //filter := bson.D{{"name", "Cloud"}}


	//1.Insert mongoDB data
	insertResult, err := collection.InsertOne(context.TODO(), bson.D{
		{Key:"name",Value:"test2"},
		{Key:"code",Value:"10-12"},
		{Key:"cusvalue",Value:"hud"},
	})

	if err != nil {
		logger.Fatal(err)
	}

    fmt.Println("Add data : ", insertResult.InsertedID)
}

func TestMongoDBFind(t *testing.T) {
	//Find mongoDB data
	mongoDBConnect()
    collection := client.Database("local").Collection("analyzer")

	findResultCursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		logger.Fatal(err)
	}

	if err = findResultCursor.All(context.TODO(), bson.D{}); err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Add data : ", findResult.ID())

	for findResultCursor.Next(context.TODO()) {
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
}
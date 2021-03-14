package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"reflect"

	//"encoding/json"

	"github.com/draftms/go_library/logging"
	"github.com/kardianos/service"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type program struct{}

var logger = logging.NewInstance()

type analyzer struct {
	name string `bson:"logName"`
	code string `bson:"logCode"`
	cusvalue string `bson:"bson:logValue"`
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	for {
		time.Sleep(1*time.Second)
		logger.Info("Windows Service Action : Service Loop" + time.Now().String())
	}

}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

var servicLoger service.Logger

func main() {
    //1. for windows service
/* 
 	svcConfig := &service.Config{
		Name:        "GoSVCTest",
		DisplayName: "Go SVC Test",
		Description: "This is a test Go service.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Fatal(err)
	}
	servicLoger, err = s.Logger(nil)
	if err != nil {
		logger.Fatal(err)
	}
	err = s.Run()
	if err != nil {
		servicLoger.Error(err)
	}
 */

/*  
    defer func() {
        err = client.Disconnect(context.TODO())

        if err != nil {
            log.Fatal(err)
        }
    }()
*/

    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/admin?safe=true")

    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil { 
        logger.Error(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        logger.Error(err)
    }

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

	//2.Find mongoDB data
	findResultCursor, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		logger.Fatal(err)
	}

	if err = findResultCursor.All(context.TODO(), bson.D{}); err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Add data : ", findResult.ID())

	for findResultCursor.Next(context.Background()) {
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

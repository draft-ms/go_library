package main

import (
	"context"
	"fmt"
	"time"

	"github.com/draftms/go_library/logging"
	"github.com/kardianos/service"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type program struct{}

var logger = logging.NewInstance()

type user struct {
/*     _id string
    email string
    password string
    name string
    role string
    hospitalid string
    status string
    hospitals string */
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

    clientOptions := options.Client().ApplyURI("mongodb://xxx:xxx@localhost:27027/admin?safe=true")

    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil { 
        logger.Error(err)
    }

    err = client.Ping(context.TODO(), nil)

    if err != nil {
        logger.Error(err)
    }

    collection := client.Database("cloudInfo").Collection("user")

    //var result user

    //filter := bson.D{{"name", "Cloud"}}

    filterCursor, err := collection.Find(context.TODO(), bson.M{"name":"Cloud"})
    if err != nil {
        logger.Fatal(err)
    }

    var userList []bson.M
    if err = filterCursor.All(context.TODO(), &userList); err != nil {
        logger.Fatal(err)
    }

    fmt.Println("Connection : ", clientOptions.AppName)
}
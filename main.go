package main

import (
	"time"

	"github.com/draftms/go_library/logging"
	"github.com/kardianos/service"
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

/*  
    defer func() {
        err = client.Disconnect(context.TODO())

        if err != nil {
            log.Fatal(err)
        }
    }()
*/
}

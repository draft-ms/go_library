package main 

import (
	"fmt"
	"github.com/draftms/go_library/configuration"
	"github.com/draftms/go_library/logging"
)

//Configuration .
var Configuration config.Configuration = config.GetConfig()
var logger = logging.NewInstance()

func init() {

}

func main() {

	//////How to use logrus
	//print log as loglevel
	fmt.Printf("loglevel : %s \n", Configuration.LOG_LEVEL)
	logger.Info("Info level log")
	logger.Warn("Warn level log")
	logger.Debug("Debug level log")
	logger.Error("Error level log")

	//to-do Fields 
	/*
	logger.WithFields(logger.Fields{
		"addField1": "field1_val",
		"addField2": "field2_val",
	}).Error("Added fields error level log")
	*/
	//////How to use logrus
}
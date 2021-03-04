package main 

import (
	"fmt"
	"github.com/draftms/library/configuration"
	"github.com/sirupsen/logrus"
)

//Configuration .
var Configuration config.Configuration = config.GetConfig()

func init() {
	//Set log level
	switch Configuration.LOG_LEVEL {
		case "DEBUG":
			logrus.SetLevel(logrus.DebugLevel)
		case "ERROR":
			logrus.SetLevel(logrus.ErrorLevel)
		case "INFO":
			logrus.SetLevel((logrus.InfoLevel))
		case "WARN":
			logrus.SetLevel((logrus.WarnLevel))
	}
}

func main() {
	//////How to use logrus
	//print log as loglevel
	fmt.Printf("loglevel : %s \n", Configuration.LOG_LEVEL)
	logrus.Info("Info level log")
	logrus.Warn("Warn level log")
	logrus.Debug("Debug level log")
	logrus.Error("Error level log")

	logrus.WithFields(logrus.Fields{
		"addField1": "field1_val",
		"addField2": "field2_val",
	}).Error("Added fields error level log")
	//////How to use logrus

}
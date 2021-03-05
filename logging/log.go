package logging

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/draftms/go_library/configuration"
)

var gLoger *logrus.Logger
var gLogEntry *logrus.Entry
var Configuration config.Configuration = config.GetConfig()

func NewInstance() *logrus.Logger {

	if gLoger != nil {
		return gLoger
	}

	gLoger = logrus.New()
	gLogEntry = logrus.NewEntry(gLoger)

	//Set log level
	switch Configuration.LOG_LEVEL {
		case "DEBUG":
			gLoger.SetLevel(logrus.DebugLevel)
		case "ERROR":
			gLoger.SetLevel(logrus.ErrorLevel)
		case "INFO":
			gLoger.SetLevel((logrus.InfoLevel))
		case "WARN":
			gLoger.SetLevel((logrus.WarnLevel))
		default:
			gLoger.SetLevel((logrus.DebugLevel))
	}
	
	//Set log format
	gLoger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	}) 	
	
	//for logfile
	filelog, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	} 
	//defer filelog.Close()	
	
	gLoger.SetOutput(filelog)

	//logrus.Fields
	//type Fields map[string]interface{}
	return gLoger
}
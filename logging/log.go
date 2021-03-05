package logging

import (
	"time"
	//"fmt"
	"github.com/draftms/go_library/configuration"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var gLoger *logrus.Logger
var Configuration config.Configuration = config.GetConfig()

func NewInstance() *logrus.Logger {

	if gLoger != nil {
		return gLoger
	}

	gLoger = logrus.New()

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
	/*
 	filelog, err := os.OpenFile("test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	} 
	//to-do 차후 defer 처리 필요
	//defer filelog.Close()	
 	*/

	writer, err := rotatelogs.New(
		//fmt.Sprintf("%s.%s", path, "%Y-%m-%d.%H:%M:%S"),
		"log.%Y%m%d.log",
		//rotatelogs.WithLinkName("."),
		rotatelogs.WithMaxAge(24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour),
		rotatelogs.WithRotationSize(1000000),
	)
	if err != nil {
		panic(err)
	}
	
	gLoger.SetOutput(writer)
	//gLoger.SetOutput(filelog)

	return gLoger
}
package logging

import (
	"time"
	"github.com/draftms/go_library/configuration"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var gLoger *logrus.Logger
var configuration config.Configuration

func NewInstance() *logrus.Logger {

	configuration = configuration.GetConfig()

	if gLoger != nil {
		return gLoger
	}

	gLoger = logrus.New()

	//Set log level
	switch configuration.LOG_LEVEL {
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

	writer, err := rotatelogs.New(
		configuration.LOG_PATH + "log.%Y%m%d.log",
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
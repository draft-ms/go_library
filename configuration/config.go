package config

import (
	"github.com/tkanos/gonfig"
	"fmt"
)

type Configuration struct {
	LOG_LEVEL string
	LOG_DIR	string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	fileName := fmt.Sprintf("./configuration/config.%s.json", env)
	gonfig.GetConf(fileName, &configuration)

	return configuration
}


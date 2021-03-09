package config

import (
	"github.com/tkanos/gonfig"
	"fmt"
	"path/filepath"
	"os"
)

type Configuration struct {
	LOG_LEVEL string
	LOG_DIR	string
	LOG_PATH string
}

func (conf Configuration)GetConfig(params ...string) Configuration {
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//gLoger.Info("Failed get filepath")
	}
 
    fileName := fmt.Sprintf(dir + "\\configuration\\config.%s.json", env)	

	gonfig.GetConf(fileName, &conf)

	return conf
}

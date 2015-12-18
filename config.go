package main

import (
	"encoding/json"
	"github.com/lonelycode/tyk-auth-proxy/tyk-api"
	"io/ioutil"
)

type Configuration struct {
	BackEnd struct {
		Name            string
		BackendSettings interface{}
	}
	TykAPISettings tyk.TykAPI
}

func loadConfig(filePath string, configStruct *Configuration) {
	configuration, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Error("Couldn't load configuration file: ", err)
		loadConfig("tap.conf", configStruct)
	} else {
		jsErr := json.Unmarshal(configuration, &configStruct)
		if jsErr != nil {
			log.Error("Couldn't unmarshal configuration: ", jsErr)
		}
	}

	log.Warning(configStruct.TykAPISettings)
}

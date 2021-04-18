package configs

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	TARGET_BITS	int
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	fileName := fmt.Sprintf("./configs/%s_config.json", env)
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		fmt.Println(err)
	}
	return configuration
}
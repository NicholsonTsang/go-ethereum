package config
import (
	"fmt"
	"github.com/tkanos/gonfig"
)
type Configuration struct {
	AUTH_CONTRACT_ADDRESS string
	CLIENT_ADDRESS string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	fileName := fmt.Sprintf("./config/%s_config.json", env)
	fmt.Println(fileName)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}
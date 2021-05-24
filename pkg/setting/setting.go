package setting

import (
	"github.com/spf13/viper"
	"log"
)

type config struct {
	App appConfig
	Mysql mysqlConfig
}
var Config = &config{}

type appConfig struct {
	Runtime string
}

type mysqlConfig struct {
	Host string
	Port int
	User string
	Password string
	DbName string
	TablePrefix string
}

func Setup() {
	//viper.SetConfigName("app")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("conf")
	viper.SetConfigFile("conf/app.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Load app yaml file error: %v", err)
	}

	viper.Unmarshal(Config)
}

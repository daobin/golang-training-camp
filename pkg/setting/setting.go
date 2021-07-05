package setting

import (
	"bytes"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
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
	confPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Get current path error: %v", err)
	}

	var buffer bytes.Buffer
	buffer.WriteString(confPath)
	buffer.WriteString("/../../conf/app.yaml")
	confPath = buffer.String()

	confPath, err = filepath.Abs(confPath)
	if err != nil {
		log.Fatalf("Get current path error: %v", err)
	}

	//viper.SetConfigName("app")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("conf")
	viper.SetConfigFile(confPath)

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Load app yaml file error: %v", err)
	}

	viper.Unmarshal(Config)
}

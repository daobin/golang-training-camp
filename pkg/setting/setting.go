package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type app struct {
	PageSize int
	RuntimeRootPath string
}
var AppSetting = &app{}

type server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}
var ServerSetting = &server{}

type database struct {
	Type string
	Host string
	User string
	Password string
	DbName string
	TablePrefix string
}
var DatabaseSetting = &database{}

type redis struct {
	Host string
	Auth string
	MaxIdle int
	MaxActive int
	IdleTimeout time.Duration
}
var RedisSetting = &redis{}

func Setup() {
	conf, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Error for load .ini: %v", err)
	}

	err = conf.Section("App").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Error for map to AppSetting: %v", err)
	}

	err = conf.Section("Server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Error for map to ServerSetting: %v", err)
	}
	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second

	err = conf.Section("Database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Error for map to DatabaseSetting: %v", err)
	}

	err = conf.Section("Redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Error for map to RedisSetting: %v", err)
	}
	RedisSetting.IdleTimeout *= time.Second
}

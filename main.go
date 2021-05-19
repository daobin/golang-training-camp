package main

import (
	"fmt"
	"github.com/daobin/golang-training-camp/pkg/setting"
	"github.com/daobin/golang-training-camp/router"
	"log"
	"net/http"
)

func main() {
	setting.Setup()

	r := router.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler: r,
		ReadTimeout: setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("%v", err)
	}
}

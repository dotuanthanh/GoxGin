package main

import (
	"api-server/api/server"
	"api-server/config"
	"api-server/internal"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	keepMainThread := make(chan struct{})
	configs, err := config.InitConfiguration()
	if err != nil {
		log.Println("InitConfiguration config fail  ", err.Error())
		close(keepMainThread)
	}
	inter, err := internal.Init(configs)
	if err != nil {
		log.Println("InitConfiguration internal fail  ", err.Error())
		close(keepMainThread)
	}
	//Framework adapter
	engine := gin.Default()

	appServer := server.NewServer(inter, configs, engine)
	go func() {
		appServer.Start()
	}()

	<-keepMainThread
}

package main

import (
	"fmt"
	"github.com/dotuanthanh/api-server/config"
	"github.com/dotuanthanh/api-server/infra/rdb"
	"github.com/dotuanthanh/api-server/infra/repository"
	"log"
	"net/http"
)

func main() {
	cfgs, err := config.Init()
	if err != nil {
		fmt.Println("Server die ....", err.Error())
		return
	}
	fmt.Println("Server start with port 8000")
	log.Fatalln("Server stop", http.ListenAndServe(":8000", handlerRequest(cfgs)))
}

func handlerRequest(cfgs *config.Server) http.Handler {
	configDB, err := rdb.Open(&cfgs.MySql)
	if err != nil {
		fmt.Println("Error when init db ...", err.Error())
	}
	repository.InitRepo(configDB)
	http.HandleFunc("/", homePage)
	return nil
}
func homePage(w http.ResponseWriter, r *http.Request) {
}

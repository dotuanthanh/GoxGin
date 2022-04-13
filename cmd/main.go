package main

import (
	"api-server/config"
	"api-server/internal/rdb"
	"api-server/internal/repository"
	"fmt"
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
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/login", auth)
	return nil
}

func admin(w http.ResponseWriter, r *http.Request) {
	//TODO authen

	switch r.Method {
	case "POST":
		path := r.URL.RawQuery
		switch path {
		case "member-create":
		case "member-update":
		case "member-delete":
		case "employee-create":
		case "employee-update":
		case "employee-delete":
		default:
		}
	case "GET":
		path := r.URL.RawQuery
		fmt.Println(path)
		switch path {
		case "members":
			w.Write([]byte("s"))
			w.WriteHeader(http.StatusOK)
		case "employees":
		}
	}
}
func auth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		path := r.URL.RawQuery
		switch path {
		case "login":

		case "sigin":

		case "change-password":

		case "forgot-password":

		}
	}
}

package config

import (
	"github.com/jinzhu/configor"
)

type Server struct {
	ServerHost string `default:"127.0.0.1" env:"APP_SERVER_DB_HOST"`
	ServerPort uint   `default:"8080" env:"APP_SERVER_DB_PORT"`
	MySql      MySql
	RootUser   RootUser
	Redis      Redis
}

func InitConfiguration() (*Server, error) {
	cfg := new(Server)
	if err := configor.Load(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

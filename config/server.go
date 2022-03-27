package config

import (
	"github.com/jinzhu/configor"
)

type Server struct {
	MySql MySql
}

func Init() (*Server, error) {
	cfg := new(Server)
	if err := configor.Load(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

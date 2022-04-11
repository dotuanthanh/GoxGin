package repository

import (
	"api-server/internal/rdb"
)

var db rdb.IDBHandler

func InitRepo(cfg rdb.IDBHandler) {
	db = cfg
}

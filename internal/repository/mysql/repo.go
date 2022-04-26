package mysql

import (
	"api-server/internal/rdb"
)

var rep rdb.IDBHandler

func InitRepo(cfg rdb.IDBHandler) {
	if rep == nil {
		rep = cfg
	}
}

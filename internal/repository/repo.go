package repository

import (
	"api-server/internal/rdb"
)

var db rdb.IRDB

func InitRepo(cfg rdb.IRDB) {
	db = cfg
}

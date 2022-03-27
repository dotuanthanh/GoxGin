package repository

import (
	"github.com/dotuanthanh/api-server/infra/rdb"
)

var db rdb.IDBHandler

func InitRepo(cfg rdb.IDBHandler) {
	db = cfg
}

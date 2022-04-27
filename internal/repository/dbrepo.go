package repository

import (
	"api-server/internal/db/rdb"
)

var DBRepository IDBRepository

func InitRepo(rdb rdb.IDBHandler) {
	if DBRepository == nil {
		DBRepository = rdb
	}
}

//func generateSql() string {
//
//	return ""
//}
//
//func placeHolder()  {
//
//}

package repository

import (
	"api-server/internal/rdb"
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

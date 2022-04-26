package repository

import (
	"api-server/internal/rdb"
)

var dbRepo IDBRepository

func InitRepo(rdb rdb.IDBHandler) {
	if dbRepo == nil {
		dbRepo = rdb
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

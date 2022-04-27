package rdb

import (
	"database/sql"
)

type IDBHandler interface {
	Close() error
	Insert(sql string, arg []interface{}) error
	Query(sql string, arg []interface{}) (*sql.Rows, error)
}

package repository

import "database/sql"

type IDBRepository interface {
	Query(sql string, args []interface{}) (*sql.Rows, error)
	Insert(sql string, args []interface{}) error
}

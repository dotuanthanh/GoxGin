package rdb

import (
	"context"
	"database/sql"
)

type IRDB interface {
	Close() error
	Insert(sql string) error
	Query(sql string) *sql.Rows
	BatchInsert(ctx context.Context, sql string) error
	Logger()
}

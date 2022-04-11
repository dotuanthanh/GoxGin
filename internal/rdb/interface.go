package rdb

type IDBHandler interface {
	Close() error
	Logger()
}

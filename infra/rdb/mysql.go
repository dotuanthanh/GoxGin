package rdb

import (
	"database/sql"
	"fmt"
	"github.com/dotuanthanh/api-server/config"
)

type mysql struct {
	client *sql.DB
}

func NewMySQL(config *config.MySql) (*mysql, error) {
	client := &mysql{}

	return client, nil
}

func (c *mysql) Close() error {
	return c.client.Close()
}

func (db *mysql) init(config *config.MySql) error {
	connectInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	client, err := sql.Open(config.DriverName, connectInfo)
	if err != nil {
		return err
	}

	client.SetMaxOpenConns(config.MaxOpenConn)
	client.SetMaxIdleConns(config.MaxIdleConn)

	if pingErr := client.Ping(); pingErr != nil {
		return pingErr
	}

	db.client = client

	return nil
}

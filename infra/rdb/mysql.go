package rdb

import (
	"database/sql"
	"fmt"
	"github.com/dotuanthanh/api-server/config"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	client *sql.DB
}

func Open(config *config.MySql) (IDBHandler, error) {
	client := &mysql{}
	err := client.init(config)
	if err != nil {
		fmt.Println("error when create instance mysql..", err.Error())
		defer func(client *mysql) {
			_ = client.Close()
		}(client)
	}
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
		fmt.Println(err.Error())
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
func (c *mysql) Logger() {
	//TODO implement me
	panic("implement me")
}
package rdb

import (
	"api-server/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	client *sql.DB
}

func NewMysql(config *config.MySql) (IDBHandler, error) {
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

//TODO do avoid sql injection
func (c *mysql) Insert(sql string, args []interface{}) error {
	exec, err := c.client.Prepare(sql)
	_, err = exec.Exec(args)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	return nil
}

//TODO do avoid sql injection
func (c *mysql) Query(sql string, args []interface{}) (*sql.Rows, error) {
	prepare, err := c.client.Prepare(sql)
	if err != nil {
		return nil, err
	}
	rows, err := prepare.Query(args)
	if err != nil {
		return nil, err
	}
	return rows, nil

}

func (c *mysql) Close() error {
	return c.client.Close()
}

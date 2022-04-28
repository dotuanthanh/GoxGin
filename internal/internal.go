package internal

import (
	"api-server/config"
	rdb2 "api-server/internal/db/rdb"
	"api-server/internal/repository"
	"api-server/internal/thirdparty/google"
)

type Internal struct {
	rdb        rdb2.IDBHandler
	thirdParty google.IThirdParty
}

func Init(configs *config.Server) (*Internal, error) {
	//TODO rdb & third-party
	db, err := rdb2.NewMysql(&configs.MySql)
	if err != nil {
		return nil, err
	}
	defer func() {
		repository.InitRepo(db)
	}()

	return &Internal{
		rdb: db,
	}, nil

}
func (i *Internal) Stop() error {
	err := i.rdb.Close()
	return err
}

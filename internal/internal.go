package internal

import (
	"api-server/config"
	"api-server/internal/rdb"
	"api-server/internal/thirdparty"
)

type Internal struct {
	rdb        rdb.IDBHandler
	thirdParty thirdparty.IThirdParty
}

func Init(configs *config.Server) (*Internal, error) {
	//TODO rdb & third-party
	db, err := rdb.NewMysql(&configs.MySql)
	if err != nil {
		return nil, err
	}
	return &Internal{
		rdb: db,
	}, nil

}

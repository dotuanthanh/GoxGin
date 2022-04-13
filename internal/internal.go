package internal

import (
	"api-server/internal/rdb"
	"api-server/internal/thirdparty"
)

type internal struct {
	rdb        rdb.IRDB
	thirdParty thirdparty.IThirdParty
}

func Init() *internal {
	//TODO rdb & third-party

	return &internal{}

}

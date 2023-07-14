package service

import (
	"github.com/switfs/switfs-block/server/dto"
	"github.com/switfs/switfs-block/server/models"
	"github.com/switfs/switfs-block/utils/mysql-rpc"
	"time"
)

type MinerId struct{}

var MinerIdAd *MinerId

func NewMinerIdService() *MinerId {
	if MinerIdAd == nil {
		lock.RLock()
		if MinerIdAd == nil {
			MinerIdAd = &MinerId{}
		}
		lock.RUnlock()
	}
	return MinerIdAd
}

func (miner MinerId) Add(addr string) (err error) {

	database := dto.NewMinerBlockTotal(mysql.RPC)
	data := models.Miner{
		MinerAddress: addr,
		MinerCreate:  time.Now(),
		MinerUpdate:  time.Now(),
	}

	return database.Create(&data)
}

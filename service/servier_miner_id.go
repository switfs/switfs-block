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
	tx := mysql.RPC.Begin()
	defer closeTx(tx, &err)

	database := dto.NewMinerBlockTotal(tx)
	data := models.Miner{
		MinerAddress: addr,
		MinerCreate:  time.Now(),
		MinerUpdate:  time.Now(),
	}

	return database.Create(&data)
}

func (miner MinerId) Del(addr string) (err error) {
	tx := mysql.RPC.Begin()
	defer closeTx(tx, &err)

	database := dto.NewMinerBlockTotal(tx)
	data := models.Miner{
		MinerAddress: addr,
		MinerCreate:  time.Now(),
		MinerUpdate:  time.Now(),
	}

	return database.Create(&data)
}

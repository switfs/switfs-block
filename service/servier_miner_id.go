package service

import "github.com/switfs/switfs-block/utils/mysql-rpc"

type MinerId struct{}

var MinerIdAd *MinerId

func NewMinerIdService() *MinerId {
	if MinerIdAd == nil {
		l.Lock()
		if MinerIdAd == nil {
			MinerIdAd = &MinerId{}
		}
		l.Unlock()
	}
	return MinerIdAd
}

func (miner MinerId) Add(addr string) (err error) {
	tx := mysql.RPC.Begin()
	defer closeTx(tx, &err)
	return nil
}

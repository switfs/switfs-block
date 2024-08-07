package service

import (
	"github.com/switfs/switfs-block/utils/mysql-rpc"
	"time"
)

func MinerId() (reust []string, err error) {
	sqlx := `SELECT miner  FROM venus_auth.miners t  WHERE t.open_mining=1`

	err = mysql.RPC.Raw(sqlx).Scan(&reust).Error
	log.Info("MinerId: ", reust)
	if err != nil {
		log.Error("myssql error:", err)
		return
	}
	return
}

func MinerUP(cid, epoch, miner, Reward string) error {
	sqlx := `UPDATE venus_miner.miner_blocks SET  cid= '` + cid + `', mine_state=1 , reward= '` + Reward + `',  consuming=1  WHERE parent_epoch=` + epoch + ` AND miner='` + miner + `' AND consuming=0`
	log.Info("sqlx: ", sqlx)
	err := mysql.RPC.Exec(sqlx).Error
	if err != nil {
		log.Error("myssql error:", err)
		return err
	}
	return nil
}

func Start() {
	go start()
}

func start() {
	for {
		st, err := MinerId()
		if err != nil {
			return
		}
		for _, v := range st {
			//Getdata("f" + v)
			err = GetPostData("f" + v)
			if err != nil {
				log.Error("GetPostData error:", err)
				continue
			}
			time.Sleep(time.Minute * 1)
		}
		time.Sleep(time.Minute * 2)
	}

}

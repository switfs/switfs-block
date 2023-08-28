package service

import "github.com/switfs/switfs-block/utils/mysql-rpc"

func MinerId() (reust []string, err error) {
	sqlx := `SELECT miner  FROM venus_auth.miners t  WHERE t.open_mining=1`

	err = mysql.RPC.Raw(sqlx).Scan(&reust).Error
	if err != nil {
		return
	}
	return
}

func MinerUP(cid, epoch, miner string) error {

	sqlx := `UPDATE venus_miner.miner_blocks SET  cid= '` + cid + `', mine_state=1 WHERE parent_epoch=` + epoch + ` AND miner='` + miner + `'`
	err := mysql.RPC.Exec(sqlx).Error
	if err != nil {
		return err
	}
	return nil
}

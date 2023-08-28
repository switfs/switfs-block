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

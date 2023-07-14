package model

import "time"

type BlockTotal struct {
	ID          uint      `gorm:"primarykey"`
	MinerId     string    `gorm:"miner_id" json:"miner_id"`
	MinerCid    string    `gorm:"miner_cid" json:"miner_cid"`
	MinerCreate time.Time `gorm:"miner_create" json:"miner_create"`
}

func (BlockTotal) TableName() string {
	return "miner_block_total"
}

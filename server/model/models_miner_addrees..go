package model

import "time"

type Miner struct {
	ID           uint      `gorm:"primarykey"`
	MinerAddress string    `gorm:"miner_address" json:"miner_address"`
	MinerCreate  time.Time `gorm:"miner_create" json:"miner_create"`
	MinerUpdate  time.Time `gorm:"miner_update" json:"miner_update"`
}

func (Miner) TableName() string {
	return "miner_addresses"
}

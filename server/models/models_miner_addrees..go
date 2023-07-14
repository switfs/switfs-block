package models

import (
	"time"
)

type Miner struct {
	ID           uint      `gorm:"primary"`
	MinerAddress string    `gorm:"miner_address" json:"miner_address"`
	MinerCreate  time.Time `gorm:"miner_create" json:"miner_create"`
	MinerUpdate  time.Time `gorm:"miner_update" json:"miner_update"`
	IsDelete     int       `gorm:"is_delete" json:"is_delete"`
}

func (Miner) TableName() string {
	return "miner_addresses"
}

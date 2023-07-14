package dto

import (
	"github.com/switfs/switfs-block/server/models"
	"gorm.io/gorm"
)

type BlockTotal struct {
	db *gorm.DB
}

func NewMinerBlockTotal(db *gorm.DB) *BlockTotal {
	return &BlockTotal{
		db: db,
	}
}

func (dto *BlockTotal) Create(block *models.Miner) error {
	log.Info("插入数据")
	return dto.Create(block)
}

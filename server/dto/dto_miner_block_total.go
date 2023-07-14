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
	return dto.Create(block)
}

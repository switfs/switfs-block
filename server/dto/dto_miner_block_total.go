package dto

import (
	"github.com/jinzhu/gorm"
	"github.com/switfs/switfs-block/server/model"
)

type BlockTotal struct {
	db *gorm.DB
}

func NewMinerBlockTotal(db *gorm.DB) *BlockTotal {
	return &BlockTotal{
		db: db,
	}
}

func (dto *BlockTotal) Create(block *model.BlockTotal) error {
	return dto.Create(block)
}

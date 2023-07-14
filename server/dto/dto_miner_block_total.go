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

	return dto.db.Create(block).Error
}

func (dto *BlockTotal) Delete(address string) error {
	return dto.db.Delete(models.Miner{}, "miner_address LIKE ?", address).Error
}

func (dto *BlockTotal) List() ([]models.Miner, error) {
	var miner []models.Miner
	err := dto.db.Find(&miner).Error
	if err != nil {
		return nil, err
	}
	return miner, nil
}

package databases

import (
	"trade_simulator/models"

	"gorm.io/gorm"
)

type historicalDatabase struct {
	db *gorm.DB
}

func NewHistoricalDatabase(db *gorm.DB) models.HistoricalDatabase {
	return &historicalDatabase{db: db}
}

func (instance *historicalDatabase) Create(historical *models.Historical) error {
	return instance.db.Create(historical).Error
}

func (instance *historicalDatabase) All() ([]models.Historical, error) {
	historicals := []models.Historical{}
	if err := instance.db.Find(&historicals).Error; err != nil {
		return nil, err
	}

	return historicals, nil
}

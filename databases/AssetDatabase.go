package databases

import (
	"trade_simulator/models"

	"gorm.io/gorm"
)

type assetDatabase struct {
	db *gorm.DB
}

func NewAssetDatabase(db *gorm.DB) models.AssetDatabase {
	return &assetDatabase{db: db}
}

func (instance *assetDatabase) Create(asset *models.Asset) error {
	return instance.db.Create(asset).Error
}

func (instance *assetDatabase) All() ([]models.Asset, error) {
	assets := []models.Asset{}
	if err := instance.db.Preload("Historicals").Find(&assets).Error; err != nil {
		return nil, err
	}

	return assets, nil
}

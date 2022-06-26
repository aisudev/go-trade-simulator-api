package services

import (
	"trade_simulator/managers"
	"trade_simulator/models"
)

type assetService struct {
	dm *managers.DatabaseManager
}

func NewAssetService(dm *managers.DatabaseManager) models.AssetService {
	return &assetService{dm: dm}
}

func (service *assetService) Create(form *models.AssetForm) error {
	asset := models.Asset{
		Name: form.Name,
	}

	return service.dm.AssetDatabase.Create(&asset)
}

func (service *assetService) All() ([]models.Asset, error) {
	return service.dm.AssetDatabase.All()
}

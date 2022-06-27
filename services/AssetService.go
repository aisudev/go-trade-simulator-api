package services

import (
	"time"
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
	asset := models.Asset{Name: form.Name}
	if err := service.dm.AssetDatabase.Create(&asset); err != nil {
		return err
	}

	createdAsset, err := service.OneByName(asset.Name)
	if err != nil {
		return err
	}

	historicals := []models.Historical{}
	for _, v := range form.Historicals {
		v.AssetID = createdAsset.ID
		v.At = time.Now()
		historicals = append(historicals, v)
	}
	return service.dm.HistoricalDatabase.BatchCreate(historicals)
}

func (service *assetService) All() ([]models.Asset, error) {
	return service.dm.AssetDatabase.All()
}

func (service *assetService) OneByID(id uint64) (*models.Asset, error) {
	return service.dm.AssetDatabase.FilterOne("id=?", id)
}

func (service *assetService) OneByName(name string) (*models.Asset, error) {
	return service.dm.AssetDatabase.FilterOne("name=?", name)
}

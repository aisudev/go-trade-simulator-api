package services

import (
	"trade_simulator/managers"
	"trade_simulator/models"
)

type historicalService struct {
	dm *managers.DatabaseManager
}

func NewHistoricalService(dm *managers.DatabaseManager) models.HistoricalService {
	return &historicalService{dm: dm}
}

package managers

import (
	"trade_simulator/models"
)

type ServiceManager struct {
	UserService        models.UserService
	TransactionService models.TransactionService
	AssetService       models.AssetService
	HistoricalService  models.HistoricalService
}

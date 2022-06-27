package services

import (
	"trade_simulator/managers"
	"trade_simulator/models"
)

type transactionService struct {
	dm *managers.DatabaseManager
}

func NewTransactionService(dm *managers.DatabaseManager) models.TransactionService {
	return &transactionService{dm: dm}
}

func (service *transactionService) Open(historical_id uint64, amount float64) (*models.Transaction, error) {
	return nil, nil
}

func (service *transactionService) Close(historical_id uint64, amount float64) (*models.Transaction, error) {
	return nil, nil
}

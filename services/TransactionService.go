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

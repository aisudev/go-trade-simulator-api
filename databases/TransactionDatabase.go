package databases

import (
	"trade_simulator/models"

	"gorm.io/gorm"
)

type transactionDatabase struct {
	db *gorm.DB
}

func NewTransactionDatabase(db *gorm.DB) models.TransactionDatabase {
	return &transactionDatabase{db: db}
}

func (instance *transactionDatabase) Create(transaction *models.Transaction) error {
	return instance.db.Create(transaction).Error
}

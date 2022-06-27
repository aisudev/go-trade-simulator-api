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

func (instance *transactionDatabase) FilterOne(orderBy, query string, args ...interface{}) (*models.Transaction, error) {
	transaction := models.Transaction{}
	exec := instance.db.Order(orderBy).Where(query, args...).Find(&transaction)
	if err := exec.Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

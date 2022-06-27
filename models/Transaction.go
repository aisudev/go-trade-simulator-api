package models

import "time"

type Transaction struct {
	ID        uint64    `json:"id" gorm:"primaryKey;not null;"`
	UserID    string    `json:"user_id" gorm:"not null;"`
	AssetCode string    `json:"asset_code" gorm:"varchar(10);not null;"`
	Status    string    `json:"status" gorm:"varchar(16);not null;"`
	Amount    float64   `json:"amount" gorm:"decimal(12, 2);not null;"`
	Price     float64   `json:"price" gorm:"decimal(12, 2);"`
	At        time.Time `json:"at"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type TransactionDatabase interface {
	Create(*Transaction) error
	FilterOne(string, string, ...interface{}) (*Transaction, error)
}

type TransactionService interface {
	/*
		@method: Open
		@requirement:
		- if never opened that user selected asset, create transaction without calculate the past transactions.
		- if opened that user selected asset, create transaction calculate the past transaction
			amount = (new_open_price * last_amount / old_open_price) + open_amount
		@params: { historical_id, amount }
	*/
	Open(string, uint64, float64) (*Transaction, error)

	/*
		@method: Close
		@requirement:
		- amount = (new_close_price * last_amount / latest_transaction_price) - close_amount
		@params: { historical_id, amount }
	*/
	Close(string, uint64, float64) (*Transaction, error)
}

type TransactionForm struct {
	HistoricalID uint64  `json:"historical_id"`
	Amount       float64 `json:"amount"`
}

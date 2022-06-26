package models

import "time"

type Transaction struct {
	ID         uint64    `json:"id" gorm:"primaryKey;not null;"`
	UserID     string    `json:"user_id" gorm:"not null;"`
	AssetCode  string    `json:"asset_code" gorm:"varchar(10);not null;"`
	Status     string    `json:"status" gorm:"varchar(16);not null;"`
	Amount     float64   `json:"amount" gorm:"decimal(12, 2);not null;"`
	OpenPrice  float64   `json:"open_price" gorm:"decimal(12, 2);"`
	OpenAt     time.Time `json:"open_at"`
	CloseAt    time.Time `json:"close_at"`
	ClosePrice float64   `json:"close_price" gorm:"decimal(12, 2);"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}

type TransactionDatabase interface {
	Create(*Transaction) error
}

type TransactionService interface {
}

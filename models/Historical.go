package models

import "time"

type Historical struct {
	ID      uint64    `json:"id" gorm:"primaryKey;not null;"`
	AssetID uint64    `json:"asset_id"`
	Open    float64   `json:"open" gorm:"decimal(12, 2);not null"`
	Close   float64   `json:"close" gorm:"decimal(12, 2);not null"`
	High    float64   `json:"high" gorm:"decimal(12, 2);not null"`
	Low     float64   `json:"low" gorm:"decimal(12, 2);not null"`
	At      time.Time `json:"at" gorm:"not null;"`
}

type HistoricalDatabase interface {
	BatchCreate([]Historical) error
}

type HistoricalService interface {
}

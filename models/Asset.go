package models

import "time"

type Asset struct {
	ID        uint64    `json:"id" gorm:"primaryKey;not null;"`
	Name      string    `json:"name" gorm:"varchar(16);unique;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Historicals []Historical `gorm:"foreignKey:AssetID"`
}

type AssetDatabase interface {
	Create(*Asset) error
	All() ([]Asset, error)
	FilterOne(string, ...interface{}) (*Asset, error)
}

type AssetService interface {
	Create(*AssetForm) error
	All() ([]Asset, error)
	OneByID(uint64) (*Asset, error)
	OneByName(string) (*Asset, error)
}

type AssetForm struct {
	Name        string       `json:"name"`
	Historicals []Historical `json:"historicals"`
}

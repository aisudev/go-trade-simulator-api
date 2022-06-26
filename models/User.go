package models

import "gorm.io/gorm"

type User struct {
	ID        string  `json:"id" gorm:"primaryKey;not null;"`
	Email     string  `json:"email" gorm:"varchar(128);unique;not null;"`
	Password  string  `json:"password" gorm:"varchar(64);not null;"`
	Name      string  `json:"name" gorm:"varchar(64);not null;"`
	Balance   float64 `json:"balance" gorm:"decimal(12, 2)"`
	NAV       float64 `json:"net_asset_value" gorm:"decimal(12, 2)"`
	DeletedAt gorm.DeletedAt

	Transactions []Transaction `gorm:"foreignKey:UserID;"`
}

type UserDatabase interface {
	Create(*User) error
	FilterOne(string, ...interface{}) (*User, error)
}

type UserService interface {
	Create(*SignUpForm) error
	// Authorization(username string, password string) (*User, error)
}

type SignInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpForm struct {
	Email    string `json:"emai"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

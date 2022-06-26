package databases

import (
	"trade_simulator/models"

	"gorm.io/gorm"
)

type userDatabase struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) models.UserDatabase {
	return &userDatabase{db: db}
}

func (instance *userDatabase) Create(user *models.User) error {
	return instance.db.Create(user).Error
}

func (instance *userDatabase) FilterOne(query string, args ...interface{}) (*models.User, error) {
	var user models.User
	if err := instance.db.Where(query, args).Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

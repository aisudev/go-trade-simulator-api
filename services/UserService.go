package services

import (
	"context"
	"trade_simulator/managers"
	"trade_simulator/models"

	"firebase.google.com/go/auth"
)

type userService struct {
	dm *managers.DatabaseManager
}

func NewUserService(dm *managers.DatabaseManager) models.UserService {
	return &userService{dm: dm}
}

func (service *userService) Create(form *models.SignUpForm) error {
	params := (&auth.UserToCreate{}).
		Email(form.Email).
		Password(form.Password)
	u, err := service.dm.Auth.CreateUser(context.Background(), params)
	if err != nil {
		return err
	}

	user := models.User{
		ID:       u.UID,
		Email:    form.Email,
		Password: form.Password,
		Name:     form.Name,
		Balance:  0,
		NAV:      0,
	}

	return service.dm.UserDatabase.Create(&user)
}

package services

import (
	"context"
	"errors"
	constantes "trade_simulator/constants"
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
		ID:      u.UID,
		Email:   form.Email,
		Name:    form.Name,
		Balance: 0,
		NAV:     0,
	}

	return service.dm.UserDatabase.Create(&user)
}

func (service *userService) OneByID(id string) (*models.User, error) {
	return service.dm.UserDatabase.FilterOne("id=?", id)
}

func (service *userService) Withdraw(id string, amount float64) (*models.User, error) {
	user, err := service.OneByID(id)
	if err != nil {
		return nil, err
	}

	if user.Balance < amount {
		return nil, errors.New("balance is not enoungh")
	}

	user.Balance -= amount
	if err := service.dm.UserDatabase.Update(user); err != nil {
		return nil, err
	}

	transaction := models.Transaction{
		UserID: user.ID,
		Status: constantes.WITHDRAW,
		Amount: amount,
	}
	if err := service.dm.TransactionDatabase.Create(&transaction); err != nil {
		return nil, err
	}

	return user, nil
}

func (service *userService) Deposit(id string, amount float64) (*models.User, error) {
	user, err := service.OneByID(id)
	if err != nil {
		return nil, err
	}

	user.Balance += amount
	if err := service.dm.UserDatabase.Update(user); err != nil {
		return nil, err
	}

	transaction := models.Transaction{
		UserID: user.ID,
		Status: constantes.DEPOSIT,
		Amount: amount,
	}
	if err := service.dm.TransactionDatabase.Create(&transaction); err != nil {
		return nil, err
	}

	return user, nil
}

package services

import (
	"errors"
	"trade_simulator/constants"
	"trade_simulator/managers"
	"trade_simulator/models"
	"trade_simulator/utils"
)

type transactionService struct {
	dm *managers.DatabaseManager
}

func NewTransactionService(dm *managers.DatabaseManager) models.TransactionService {
	return &transactionService{dm: dm}
}

func (service *transactionService) Open(user_id string, historical_id uint64, amount float64) (*models.Transaction, error) {
	user, err := service.dm.UserDatabase.FilterOne("id=?", user_id)
	if err != nil {
		return nil, err
	}

	if user.Balance < amount {
		return nil, errors.New("balance is not enough")
	}

	historical, err := service.dm.HistoricalDatabase.FilterOne("id=?", historical_id)
	if err != nil {
		return nil, err
	}

	asset, err := service.dm.AssetDatabase.FilterOne(false, "id=?", historical.AssetID)
	if err != nil {
		return nil, err
	}

	tx, err := service.dm.TransactionDatabase.FilterOne("at desc", "user_id = ? AND asset_code = ? AND status IN ?", user_id, asset.Name, []string{constants.CLOSE, constants.OPEN})
	if err != nil {
		return nil, err
	} else if historical.At.Sub(tx.At).Seconds() < 0 {
		return nil, errors.New("invalid transaction")
	}

	newTx := models.Transaction{
		UserID:    user_id,
		AssetCode: asset.Name,
		Status:    constants.OPEN,
		Price:     historical.Open,
		At:        *historical.At,
	}

	if tx.ID != 0 {
		newTx.Amount = utils.CalculateOpenPosition(tx.Price, tx.Amount, historical.Open, amount)
	} else {
		newTx.Amount = amount
	}

	if err = service.dm.TransactionDatabase.Create(&newTx); err != nil {
		return nil, err
	}

	user.Balance -= amount
	if err = service.dm.UserDatabase.Update(user); err != nil {
		return nil, err
	}

	return &newTx, nil
}

func (service *transactionService) Close(user_id string, historical_id uint64, amount float64) (*models.Transaction, error) {
	user, err := service.dm.UserDatabase.FilterOne("id=?", user_id)
	if err != nil {
		return nil, err
	}

	historical, err := service.dm.HistoricalDatabase.FilterOne("id=?", historical_id)
	if err != nil {
		return nil, err
	}

	asset, err := service.dm.AssetDatabase.FilterOne(false, "id=?", historical.AssetID)
	if err != nil {
		return nil, err
	}

	tx, err := service.dm.TransactionDatabase.FilterOne("at desc", "user_id = ? AND asset_code = ? AND status IN ?", user_id, asset.Name, []string{constants.CLOSE, constants.OPEN})
	if err != nil {
		return nil, err
	} else if historical.At.Sub(tx.At).Seconds() < 0 {
		return nil, errors.New("invalid transaction")
	} else if tx.Amount < amount {
		return nil, errors.New("asset value is not enough")
	}

	newTx := models.Transaction{
		UserID:    user_id,
		AssetCode: asset.Name,
		Status:    constants.CLOSE,
		Price:     historical.Open,
		At:        *historical.At,
	}

	if tx.ID == 0 {
		return nil, errors.New("none transaction")
	}
	newTx.Amount = utils.CalculateClosePosition(tx.Price, tx.Amount, historical.Close, amount)

	if err = service.dm.TransactionDatabase.Create(&newTx); err != nil {
		return nil, err
	}

	user.Balance += amount
	if err = service.dm.UserDatabase.Update(user); err != nil {
		return nil, err
	}

	return &newTx, nil

}

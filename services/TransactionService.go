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
	// query user data
	user, err := service.dm.UserDatabase.FilterOne("id=?", user_id)
	if err != nil {
		return nil, err
	}

	// verify user balance
	if user.Balance < amount {
		return nil, errors.New("balance is not enough")
	}

	// query historical data
	historical, err := service.dm.HistoricalDatabase.FilterOne("id=?", historical_id)
	if err != nil {
		return nil, err
	}

	// query asset data
	asset, err := service.dm.AssetDatabase.FilterOne(false, "id=?", historical.AssetID)
	if err != nil {
		return nil, err
	}

	// query latest transaction condition by user_id, asset_code and status
	tx, err := service.dm.TransactionDatabase.FilterOne("at desc", "user_id = ? AND asset_code = ? AND status IN ?", user_id, asset.Name, []string{constants.CLOSE, constants.OPEN})
	if err != nil {
		return nil, err
	} else if historical.At.Sub(tx.At).Seconds() < 0 {
		return nil, errors.New("invalid transaction")
	}

	// draft new transaction
	newTx := models.Transaction{
		UserID:    user_id,
		AssetCode: asset.Name,
		Status:    constants.OPEN,
		Price:     historical.Open,
		At:        *historical.At,
	}

	// calculate new amount in new transaction
	if tx.ID != 0 {
		newTx.Amount = utils.CalculateOpenPosition(tx.Price, tx.Amount, historical.Open, amount)
	} else {
		newTx.Amount = amount
	}

	// create new transaction
	if err = service.dm.TransactionDatabase.Create(&newTx); err != nil {
		return nil, err
	}

	// update user balance
	user.Balance -= amount
	if err = service.dm.UserDatabase.Update(user); err != nil {
		return nil, err
	}

	return &newTx, nil
}

func (service *transactionService) Close(user_id string, historical_id uint64, amount float64) (*models.Transaction, error) {
	// query user data
	user, err := service.dm.UserDatabase.FilterOne("id=?", user_id)
	if err != nil {
		return nil, err
	}

	// query historical data
	historical, err := service.dm.HistoricalDatabase.FilterOne("id=?", historical_id)
	if err != nil {
		return nil, err
	}

	// query asset data
	asset, err := service.dm.AssetDatabase.FilterOne(false, "id=?", historical.AssetID)
	if err != nil {
		return nil, err
	}

	// query latest transaction condition by user_id, asset_code and status
	tx, err := service.dm.TransactionDatabase.FilterOne("at desc", "user_id = ? AND asset_code = ? AND status IN ?", user_id, asset.Name, []string{constants.CLOSE, constants.OPEN})
	if err != nil {
		return nil, err
	} else if tx.ID == 0 {
		return nil, errors.New("none transaction")
	} else if historical.At.Sub(tx.At).Seconds() < 0 {
		return nil, errors.New("invalid transaction")
	} else if tx.Amount < amount {
		return nil, errors.New("asset value is not enough")
	}

	// draft new transaction
	newTx := models.Transaction{
		UserID:    user_id,
		AssetCode: asset.Name,
		Status:    constants.CLOSE,
		Price:     historical.Close,
		At:        *historical.At,
	}
	newTx.Amount = utils.CalculateClosePosition(tx.Price, tx.Amount, historical.Close, amount)

	// create new transaction
	if err = service.dm.TransactionDatabase.Create(&newTx); err != nil {
		return nil, err
	}

	// update user balance
	user.Balance += amount
	if err = service.dm.UserDatabase.Update(user); err != nil {
		return nil, err
	}

	return &newTx, nil
}

func (service *transactionService) CloseAll(user_id string, historical_id uint64) (*models.Transaction, error) {
	// query user data
	user, err := service.dm.UserDatabase.FilterOne("id=?", user_id)
	if err != nil {
		return nil, err
	}

	// query historical data
	historical, err := service.dm.HistoricalDatabase.FilterOne("id=?", historical_id)
	if err != nil {
		return nil, err
	}

	// query asset data
	asset, err := service.dm.AssetDatabase.FilterOne(false, "id=?", historical.AssetID)
	if err != nil {
		return nil, err
	}

	// query latest transaction condition by user_id, asset_code and status
	tx, err := service.dm.TransactionDatabase.FilterOne("at desc", "user_id = ? AND asset_code = ? AND status IN ?", user_id, asset.Name, []string{constants.CLOSE, constants.OPEN})
	if err != nil {
		return nil, err
	} else if tx.ID == 0 {
		return nil, errors.New("none transaction")
	} else if historical.At.Sub(tx.At).Seconds() < 0 {
		return nil, errors.New("invalid transaction")
	}

	// draft new transaction
	newTx := models.Transaction{
		UserID:    user_id,
		AssetCode: asset.Name,
		Status:    constants.CLOSE,
		Price:     historical.Close,
		Amount:    0,
		At:        *historical.At,
	}
	if err = service.dm.TransactionDatabase.Create(&newTx); err != nil {
		return nil, err
	}

	user.Balance += utils.ClosePosition(tx.Price, tx.Amount, historical.Close)
	if err = service.dm.UserDatabase.Update(user); err != nil {
		return nil, err
	}

	return &newTx, nil
}

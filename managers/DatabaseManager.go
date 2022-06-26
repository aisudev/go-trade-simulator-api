package managers

import (
	"trade_simulator/models"

	"firebase.google.com/go/auth"
)

type DatabaseManager struct {
	Auth                *auth.Client
	UserDatabase        models.UserDatabase
	TransactionDatabase models.TransactionDatabase
	AssetDatabase       models.AssetDatabase
	HistoricalDatabase  models.HistoricalDatabase
}

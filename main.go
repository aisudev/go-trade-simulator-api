package main

import (
	"trade_simulator/configs"
	"trade_simulator/controllers"
	"trade_simulator/databases"
	"trade_simulator/managers"
	"trade_simulator/middlewares"
	"trade_simulator/models"
	"trade_simulator/services"

	"firebase.google.com/go/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

var (
	Auth *auth.Client
	DB   *gorm.DB
	DM   *managers.DatabaseManager
	SM   *managers.ServiceManager
)

func init() {
	DB = configs.DatabaseConnection()
	Auth = configs.FirebaseAuthSetup()
	AutoMigration()

	DM = &managers.DatabaseManager{
		Auth:                Auth,
		UserDatabase:        databases.NewUserDatabase(DB),
		TransactionDatabase: databases.NewTransactionDatabase(DB),
		AssetDatabase:       databases.NewAssetDatabase(DB),
		HistoricalDatabase:  databases.NewHistoricalDatabase(DB),
	}

	SM = &managers.ServiceManager{
		UserService:        services.NewUserService(DM),
		TransactionService: services.NewTransactionService(DM),
		AssetService:       services.NewAssetService(DM),
		HistoricalService:  services.NewHistoricalService(DM),
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	auth := e.Group("/auth")
	controllers.NewAuthController(auth, SM)

	user := e.Group("/user")
	user.Use(middlewares.AuthMiddleware(Auth))
	controllers.NewUserController(user, SM)

	asset := e.Group("/asset")
	asset.Use(middlewares.AuthMiddleware(Auth))
	controllers.NewAssetController(asset, SM)

	transaction := e.Group("/tx")
	transaction.Use(middlewares.AuthMiddleware(Auth))
	controllers.NewTransactionController(transaction, SM)

	e.Logger.Fatal(e.Start(":5000"))
}

func AutoMigration() {
	DB.AutoMigrate(
		&models.User{},
		&models.Transaction{},
		&models.Asset{},
		&models.Historical{},
	)
}

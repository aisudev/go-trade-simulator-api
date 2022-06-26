package configs

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		"127.0.0.1",
		"5432",
		"root",
		"abcd1234",
		"tradesim",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	return db
}

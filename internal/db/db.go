package db

import (
	"api/shorturl/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func ConnectDb(config *models.Config) *Db {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic("Error connect Postgres Db")
	}
	//Create Migrate
	db.AutoMigrate(&models.Link{}, &models.User{})

	return &Db{
		db,
	}
}

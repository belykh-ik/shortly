package service

import (
	"api/shorturl/internal/db"
	"api/shorturl/internal/models"
)

type UserRepository struct {
	db *db.Db
}

func (db *UserRepository) Create(user *models.User) *models.User {
	db.db.Create(user)
	return user
}

func (db *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	res := db.db.First(&user, "Email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

package service

import (
	"api/shorturl/internal/db"
	"api/shorturl/internal/models"
	"errors"
)

type UserRepository struct {
	db *db.Db
}

func NewUserRepository(db *db.Db) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (db *UserRepository) Create(reg *models.RegisterRequest) (*models.User, error) {
	userExist, _ := db.FindByEmail(reg.Email)
	if userExist != nil {
		return nil, errors.New("EXIST")
	}
	user := &models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: reg.Password,
	}
	db.db.Create(user)
	return user, nil
}

func (db *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	res := db.db.First(&user, "Email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

package service

import (
	"api/shorturl/internal/db"
	"api/shorturl/internal/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("ERROR_ADD_PASSWORD")
	}
	user := &models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: string(hashedPassword),
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

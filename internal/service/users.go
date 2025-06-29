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

// Login User
func (db *UserRepository) LoginUser(data *models.LoginRequest) (string, error) {
	user, err := db.findByEmail(data.Email)
	if err != nil {
		return "", errors.New("USER_OR_PASSWORD_ERROR")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return "", errors.New("USER_OR_PASSWORD_ERROR")
	}
	return user.Email, nil
}

// Create New User
func (db *UserRepository) CreateUser(reg *models.RegisterRequest) (string, error) {
	userExist, _ := db.findByEmail(reg.Email)
	if userExist != nil {
		return "", errors.New("EXIST")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reg.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("ERROR_ADD_PASSWORD")
	}
	user := &models.User{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: string(hashedPassword),
	}
	db.db.Create(user)
	return user.Email, nil
}

// Search user by Email
func (db *UserRepository) findByEmail(email string) (*models.User, error) {
	var user models.User
	res := db.db.First(&user, "Email = ?", email)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

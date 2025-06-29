package models

import "gorm.io/gorm"

type Config struct {
	DSN    string
	Secret string
}

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" validate:"unique"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"index" validate:"unique"`
	Password string
}

type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type Url struct {
	Url string `json:"url" validate:"url"`
}

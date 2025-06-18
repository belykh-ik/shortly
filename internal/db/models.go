package db

import (
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: generateHash(15),
	}
}

// Generate Hash for Url
func generateHash(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rune(rand.Intn(len([]rune("qwertyuioplkjhgfdsazxcvbnmqwQWERTYUIOPASDFGHJKLZXCVBNM123456789"))))
	}
	return string(b)
}

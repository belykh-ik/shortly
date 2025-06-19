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
	r := []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	b := make([]rune, n)
	for i := range b {
		b[i] = r[(rand.Intn(len(r)))]
	}
	return string(b)
}

package db

import (
	"math/rand"

	"gorm.io/gorm"
)

var (
	elementsForHash = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	existsHash      = (make(map[string][]rune, 0))
)

type Link struct {
	gorm.Model
	Url  string `json:"url"`
	Hash string `json:"hash" validate:"unique"`
}

func NewLink(url string) *Link {
	return &Link{
		Url:  url,
		Hash: generateHash(30),
	}
}

// Generate Hash for Url
func generateHash(n int) string {
	newHash := make([]rune, n)
	for {
		for i := range newHash {
			newHash[i] = elementsForHash[(rand.Intn(len(elementsForHash)))]
		}
		if _, exist := existsHash[string(newHash)]; !exist {
			existsHash[string(newHash)] = newHash
			return string(newHash)
		}
		newHash = newHash[:0]
		continue
	}
}

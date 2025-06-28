package service

import (
	"api/shorturl/internal/models"
	"math/rand"
)

var (
	elementsForHash = []rune("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	existsHash      = (make(map[string][]rune, 0))
)

func NewLink(url string) *models.Link {
	return &models.Link{
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

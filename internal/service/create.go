package service

import (
	"api/shorturl/internal/db"
	"api/shorturl/internal/models"
)

type LinkDeps struct {
	Database *db.Db
}

func NewLink(db *db.Db) *LinkDeps {
	return &LinkDeps{
		Database: db,
	}
}

func (l LinkDeps) LinkCreate(link *models.Link) {
	linkStruct := db.NewLink(link.Url)
	l.Database.Create(linkStruct)
}

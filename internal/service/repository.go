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

func (l LinkDeps) LinkUpdate(link *models.Link, id *uint64) {
	linkStruct := db.NewLink(link.Url)
	l.Database.Where("id = ?", id).Updates(linkStruct)
}

func (l LinkDeps) LinkGet(hash string) *db.Link {
	var originalLink *db.Link
	l.Database.First(&originalLink, "hash = ?", hash)
	return originalLink
}

func (l LinkDeps) LinkDelete(id *uint64) error {
	tx := l.Database.First(&db.Link{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	l.Database.Delete(&db.Link{}, uint(*id))
	return nil
}

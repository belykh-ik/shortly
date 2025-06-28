package service

import (
	"api/shorturl/internal/db"
	"api/shorturl/internal/models"
	"errors"
)

type LinkDeps struct {
	Database *db.Db
}

func NewLinkDeps(db *db.Db) *LinkDeps {
	return &LinkDeps{
		Database: db,
	}
}

func (l LinkDeps) LinkCreate(link *models.Url) {
	linkStruct := NewLink(link.Url)
	l.Database.Create(linkStruct)
}

func (l LinkDeps) LinkUpdate(link *models.Url, id *uint64) {
	linkStruct := NewLink(link.Url)
	l.Database.Where("id = ?", id).Updates(linkStruct)
}

func (l LinkDeps) LinkGet(hash string) (*models.Link, error) {
	var originalLink *models.Link
	l.Database.First(&originalLink, "hash = ?", hash)
	if originalLink.Url == "" {
		return nil, errors.New("URL_NOT_FOUND")
	}
	return originalLink, nil
}

func (l LinkDeps) LinkDelete(id *uint64) error {
	tx := l.Database.First(&models.Link{}, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	l.Database.Delete(&models.Link{}, uint(*id))
	return nil
}

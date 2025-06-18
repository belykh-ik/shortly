package service

import "api/shorturl/internal/db"

type LinkDeps struct {
	Database *db.Db
}

func NewLink(db *db.Db) *LinkDeps {
	return &LinkDeps{
		Database: db,
	}
}

func (link LinkDeps) Create() {

}

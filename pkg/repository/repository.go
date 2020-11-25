package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sashindionicus/DBDocument"
)

type Documents interface {
	Get(id int) (DBDocument.Document, error)
	Create(DBDocument.Document) (int, error)
	Delete(id int) error
}

type Repositories struct {
	Documents
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		NewDocumentPostgres(db),
	}
}

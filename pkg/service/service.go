package service

import (
	"github.com/sashindionicus/DBDocument"
	"github.com/sashindionicus/DBDocument/pkg/repository"
)

type Documents interface {
	Get(id int) (DBDocument.Document, error)
	Create(DBDocument.Document) (int, error)
	Delete(id int) error
}

type Services struct {
	Documents
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Documents: NewDocumentsService(repos.Documents),
	}
}

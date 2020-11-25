package service

import (
	"github.com/sashindionicus/DBDocument"
	"github.com/sashindionicus/DBDocument/pkg/repository"
)

type DocumentsService struct {
	repo repository.Documents
}

func NewDocumentsService(repo repository.Documents) *DocumentsService {
	return &DocumentsService{
		repo: repo,
	}
}

func (s *DocumentsService) Get(id int) (DBDocument.Document, error) {
	return s.repo.Get(id)
}

func (s *DocumentsService) Create(doc DBDocument.Document) (int, error) {
	return s.repo.Create(doc)
}

func (s *DocumentsService) Delete(id int) error {
	return s.repo.Delete(id)
}

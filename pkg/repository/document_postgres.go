package repository

import (
	"fmt"
	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"
	"github.com/sashindionicus/DBDocument"
)

type DocumentPostgres struct {
	db *sqlx.DB
}

func NewDocumentPostgres(db *sqlx.DB) *DocumentPostgres {
	return &DocumentPostgres{db: db}
}

func (dp *DocumentPostgres) Get(id int) (DBDocument.Document, error) {
	var document DBDocument.Document

	var authorId null.Int
	row := dp.db.QueryRow(fmt.Sprintf("SELECT id, title, author_id FROM %s WHERE id = $1", documentsTable), id)
	err := row.Scan(&document.ID, &document.Title, &authorId)
	if err != nil {
		return document, err
	}

	if authorId.Valid {
		document.Author = &DBDocument.Author{}
		err := dp.db.Get(document.Author, fmt.Sprintf("SELECT first_name, last_name FROM %s WHERE id = $1", authorsTable),
			authorId.Int64)
		if err != nil {
			return document, err
		}
	}

	return document, nil
}

func (dp *DocumentPostgres) Create(doc DBDocument.Document) (int, error) {
	if doc.Author != nil {
		return dp.createDocumentWithAuthor(doc)
	}

	return dp.createDocumentWithoutAuthor(doc)
}

func (dp *DocumentPostgres) Delete(id int) (err error) {
	_, err = dp.db.Exec(fmt.Sprintf("DELETE FROM %s WHERE id=$1", documentsTable), id)
	if err != nil {
		return err
	}
	return nil
}

func (dp *DocumentPostgres) createDocumentWithAuthor(doc DBDocument.Document) (int, error) {
	tx, err := dp.db.Begin()
	if err != nil {
		return 0, nil
	}

	var authorId int
	row := tx.QueryRow(fmt.Sprintf("INSERT INTO %s (first_name, last_name) values ($1, $2) RETURNING id", authorsTable),
		doc.Author.Firstname, doc.Author.Lastname)
	err = row.Scan(&authorId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var docId int
	row = tx.QueryRow(fmt.Sprintf("INSERT INTO %s (title, author_id) values ($1, $2) RETURNING id", documentsTable),
		doc.Title, authorId)
	err = row.Scan(&docId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return docId, tx.Commit()
}

func (dp *DocumentPostgres) createDocumentWithoutAuthor(doc DBDocument.Document) (int, error) {
	var docId int
	row := dp.db.QueryRow(fmt.Sprintf("INSERT INTO %s (title) values ($1) RETURNING id", documentsTable),
		doc.Title)
	err := row.Scan(&docId)

	return docId, err
}

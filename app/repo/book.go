package repo

import (
	"ysqlbm/db"
	"ysqlbm/schema"

	"github.com/google/uuid"
)

type BookRepo interface {
	Create(book *schema.Book) error
	Update(book *schema.Book) error
	GetById(id uuid.UUID) (*schema.Book, error)
	List(offset, limit int, genreId, authorId string) ([]schema.Book, error)
	Delete(id uuid.UUID) error
}

var _ BookRepo = new(bookRepo)

type bookRepo struct {
	db *db.CDB
}

func NewBookDAL(db *db.CDB) BookRepo {
	return &bookRepo{
		db: db,
	}
}

func (a *bookRepo) Create(book *schema.Book) error {
	return a.db.Create(book).Error
}

func (a *bookRepo) GetById(id uuid.UUID) (*schema.Book, error) {
	var book schema.Book
	if err := a.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (a *bookRepo) Update(book *schema.Book) error {
	return a.db.Save(book).Error
}

func (a *bookRepo) List(offset, limit int, genreId, authorId string) ([]schema.Book, error) {
	var books []schema.Book

	dbQuery := a.db.Offset(offset).Limit(limit)

	if authorId != "" {
		dbQuery = dbQuery.Where("author_id = ?", authorId)
	}

	if genreId != "" {
		dbQuery = dbQuery.Where("genre_id = ?", genreId)
	}

	err := dbQuery.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (a *bookRepo) Delete(id uuid.UUID) error {
	return a.db.Delete(&schema.Book{}, id).Error
}

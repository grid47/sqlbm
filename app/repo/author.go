package repo

import (
	"ysqlbm/db"
	"ysqlbm/schema"

	"github.com/google/uuid"
)

type AuthorRepo interface {
	Create(author *schema.Author) error
	Update(author *schema.Author) error
	GetById(id uuid.UUID) (*schema.Author, error)
	List(offset, limit int) ([]schema.Author, error)
	Delete(id uuid.UUID) error
}

var _ AuthorRepo = new(authorRepo)

type authorRepo struct {
	db *db.CDB
}

func NewAuthorDAL(db *db.CDB) AuthorRepo {
	return &authorRepo{
		db: db,
	}
}

func (a *authorRepo) Create(author *schema.Author) error {
	return a.db.Create(author).Error
}

func (a *authorRepo) GetById(id uuid.UUID) (*schema.Author, error) {
	var author schema.Author
	if err := a.db.First(&author, id).Error; err != nil {
		return nil, err
	}
	return &author, nil
}

func (a *authorRepo) Update(author *schema.Author) error {
	return a.db.Save(author).Error
}

func (a *authorRepo) List(offset, limit int) ([]schema.Author, error) {
	var authors []schema.Author
	if err := a.db.Offset(offset).Limit(limit).Find(&authors).Error; err != nil {
		return nil, err
	}
	return authors, nil
}

func (a *authorRepo) Delete(id uuid.UUID) error {
	return a.db.Delete(&schema.Author{}, id).Error
}

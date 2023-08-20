package repo

import (
	"ysqlbm/db"
	"ysqlbm/schema"

	"github.com/google/uuid"
)

type GenreRepo interface {
	Create(genre *schema.Genre) error
	Update(genre *schema.Genre) error
	GetById(id uuid.UUID) (*schema.Genre, error)
	List(offset, limit int) ([]schema.Genre, error)
	Delete(id uuid.UUID) error
}

var _ GenreRepo = new(genreRepo)

type genreRepo struct {
	db *db.CDB
}

func NewGenreDAL(db *db.CDB) GenreRepo {
	return &genreRepo{
		db: db,
	}
}

func (a *genreRepo) Create(genre *schema.Genre) error {
	return a.db.Create(genre).Error
}

func (a *genreRepo) GetById(id uuid.UUID) (*schema.Genre, error) {
	var genre schema.Genre
	if err := a.db.First(&genre, id).Error; err != nil {
		return nil, err
	}
	return &genre, nil
}

func (a *genreRepo) Update(genre *schema.Genre) error {
	return a.db.Save(genre).Error
}

func (a *genreRepo) List(offset, limit int) ([]schema.Genre, error) {
	var genres []schema.Genre
	if err := a.db.Offset(offset).Limit(limit).Find(&genres).Error; err != nil {
		return nil, err
	}
	return genres, nil
}

func (a *genreRepo) Delete(id uuid.UUID) error {
	return a.db.Delete(&schema.Genre{}, id).Error
}

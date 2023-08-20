package schema

import "github.com/google/uuid"

type Author struct {
	Id     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name   string
	Detail string
}

type Book struct {
	Id       uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string
	Detail   string
	Author   string
	AuthorId uuid.UUID
	Genre    string
	GenreId  uuid.UUID
}

type Genre struct {
	Id     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name   string
	Detail string
}

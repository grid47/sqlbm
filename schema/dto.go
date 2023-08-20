package schema

import "github.com/google/uuid"

type CreateAuthor struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type UpdateAuthor struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Detail string    `json:"detail"`
}

type CreateGenre struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
}

type UpdateGenre struct {
	Id     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Detail string    `json:"detail"`
}

type CreateBook struct {
	Name     string    `json:"name"`
	Detail   string    `json:"detail"`
	AuthorId uuid.UUID `json:"authorId"`
	GenreId  uuid.UUID `json:"genreId"`
}

type UpdateBook struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Detail   string    `json:"detail"`
	AuthorId uuid.UUID `json:"authorId"`
	GenreId  uuid.UUID `json:"genreId"`
}

const DefaultLimit = 10

type Page struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type ListBooks struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	AuthorId string `json:"authorId"`
	GenreId  string `json:"genreId"`
}

type Response struct {
	Error string `json:"error"`
	Data  any    `json:"data"`
}

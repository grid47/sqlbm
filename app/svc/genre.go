package svc

import (
	"fmt"
	"ysqlbm/app/repo"
	"ysqlbm/schema"

	"github.com/google/uuid"
)

type GenreSvc interface {
	Create(req schema.CreateGenre) (*schema.Genre, error)
	Update(req schema.UpdateGenre) error
	Get(id uuid.UUID) (*schema.Genre, error)
	List(page schema.Page) ([]schema.Genre, error)
}

var _ GenreSvc = new(genreSvc)

func NewGenreSvc(genreDal repo.GenreRepo) GenreSvc {
	return &genreSvc{
		genrerepo: genreDal,
	}
}

type genreSvc struct {
	genrerepo repo.GenreRepo
}

func (a *genreSvc) Create(req schema.CreateGenre) (*schema.Genre, error) {
	genre := &schema.Genre{
		Id:     uuid.New(),
		Name:   req.Name,
		Detail: req.Detail,
	}

	err := a.genrerepo.Create(genre)
	if err != nil {
		fmt.Printf("genre creation failed %v \n", req)
		return nil, err
	}

	fmt.Printf("genre created %v \n", req)
	return genre, nil
}

func (a *genreSvc) Update(req schema.UpdateGenre) error {
	genre, err := a.genrerepo.GetById(req.Id)
	if err != nil {
		fmt.Printf("genre not found for update %v \n", req)
		return err
	}

	genre.Name = req.Name
	genre.Detail = req.Detail

	err = a.genrerepo.Update(genre)
	if err != nil {
		fmt.Printf("genre update failed %v \n", req)
		return err
	}

	fmt.Printf("genre updated %v \n", req)
	return nil
}

func (a *genreSvc) Get(id uuid.UUID) (*schema.Genre, error) {
	genre, err := a.genrerepo.GetById(id)
	if err != nil {
		fmt.Printf("get genre failed, id - %v \n", id)
		return nil, err
	}

	fmt.Printf("retrieved genre %v \n", genre)
	return genre, nil
}

func (a *genreSvc) List(page schema.Page) ([]schema.Genre, error) {

	if page.Limit > schema.DefaultLimit {
		page.Limit = schema.DefaultLimit
	}

	genres, err := a.genrerepo.List(page.Page*page.Limit, page.Limit)
	if err != nil {
		fmt.Println("failed to fetch genres")
		return nil, err
	}

	fmt.Println("retrieved genres", genres)
	return genres, nil
}

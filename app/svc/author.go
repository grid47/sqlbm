package svc

import (
	"fmt"
	"ysqlbm/app/repo"
	"ysqlbm/schema"

	"github.com/google/uuid"
)

type AuthorSvc interface {
	Create(req schema.CreateAuthor) (*schema.Author, error)
	Update(req schema.UpdateAuthor) error
	Get(id uuid.UUID) (*schema.Author, error)
	List(page schema.Page) ([]schema.Author, error)
	Delete(id uuid.UUID) error
}

var _ AuthorSvc = new(authorSvc)

func NewAuthorSvc(authorDal repo.AuthorRepo) AuthorSvc {
	return &authorSvc{
		authorrepo: authorDal,
	}
}

type authorSvc struct {
	authorrepo repo.AuthorRepo
}

func (a *authorSvc) Create(req schema.CreateAuthor) (*schema.Author, error) {
	author := &schema.Author{
		Id:     uuid.New(),
		Name:   req.Name,
		Detail: req.Detail,
	}

	err := a.authorrepo.Create(author)
	if err != nil {
		fmt.Printf("author creation failed %v \n", req)
		return nil, err
	}

	fmt.Printf("author created %v \n", req)
	return author, nil
}

func (a *authorSvc) Update(req schema.UpdateAuthor) error {
	author, err := a.authorrepo.GetById(req.Id)
	if err != nil {
		fmt.Printf("author not found for update %v \n", req)
		return err
	}

	author.Name = req.Name
	author.Detail = req.Detail

	err = a.authorrepo.Update(author)
	if err != nil {
		fmt.Printf("author update failed %v \n", req)
		return err
	}

	fmt.Printf("author updated %v \n", req)
	return nil
}

func (a *authorSvc) Get(id uuid.UUID) (*schema.Author, error) {
	author, err := a.authorrepo.GetById(id)
	if err != nil {
		fmt.Printf("get author failed, id - %v \n", id)
		return nil, err
	}

	fmt.Printf("retrieved author %v \n", author)
	return author, nil
}

func (a *authorSvc) List(page schema.Page) ([]schema.Author, error) {

	if page.Limit > schema.DefaultLimit {
		page.Limit = schema.DefaultLimit
	}

	authors, err := a.authorrepo.List(page.Page*page.Limit, page.Limit)
	if err != nil {
		fmt.Println("failed to fetch authors")
		return nil, err
	}

	fmt.Println("retrieved authors", authors)
	return authors, nil
}

func (a *authorSvc) Delete(id uuid.UUID) error {
	err := a.authorrepo.Delete(id)
	if err != nil {
		fmt.Printf("author deletion failed, id - %v \n", id)
		return err
	}

	fmt.Printf("author deleted, id - %v \n", id)
	return nil
}

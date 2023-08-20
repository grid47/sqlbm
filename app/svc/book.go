package svc

import (
	"fmt"
	"ysqlbm/app/repo"
	"ysqlbm/schema"

	"github.com/google/uuid"
)

type BookSvc interface {
	Create(req schema.CreateBook) (*schema.Book, error)
	Update(req schema.UpdateBook) error
	Get(id uuid.UUID) (*schema.Book, error)
	List(page schema.ListBooks) ([]schema.Book, error)
	Delete(id uuid.UUID) error
}

var _ BookSvc = new(bookSvc)

func NewBookSvc(bookDal repo.BookRepo,
	gen repo.GenreRepo,
	author repo.AuthorRepo) BookSvc {
	return &bookSvc{
		bookrepo:   bookDal,
		genrerepo:  gen,
		authorrepo: author,
	}
}

type bookSvc struct {
	bookrepo   repo.BookRepo
	genrerepo  repo.GenreRepo
	authorrepo repo.AuthorRepo
}

func (a *bookSvc) Create(req schema.CreateBook) (*schema.Book, error) {

	author, err := a.authorrepo.GetById(req.AuthorId)
	if err != nil {
		fmt.Printf("author not found, book creation failed %v \n", req)
		return nil, err
	}

	genre, err := a.genrerepo.GetById(req.GenreId)
	if err != nil {
		fmt.Printf("genre not found, book creation failed %v \n", req)
		return nil, err
	}

	book := &schema.Book{
		Id:       uuid.New(),
		Name:     req.Name,
		Detail:   req.Detail,
		Author:   author.Name,
		AuthorId: author.Id,
		Genre:    genre.Name,
		GenreId:  genre.Id,
	}

	err = a.bookrepo.Create(book)
	if err != nil {
		fmt.Printf("book creation failed %v \n", req)
		return nil, err
	}

	fmt.Printf("book created %v \n", req)
	return book, nil
}

func (a *bookSvc) Update(req schema.UpdateBook) error {
	book, err := a.bookrepo.GetById(req.Id)
	if err != nil {
		fmt.Printf("book not found for update %v \n", req)
		return err
	}

	author, err := a.authorrepo.GetById(req.AuthorId)
	if err != nil {
		fmt.Printf("author not found, book update failed %v \n", req)
		return err
	}

	genre, err := a.genrerepo.GetById(req.GenreId)
	if err != nil {
		fmt.Printf("genre not found, book update failed %v \n", req)
		return err
	}

	book.Name = req.Name
	book.Detail = req.Detail
	book.Author = author.Name
	book.AuthorId = author.Id
	book.Genre = genre.Name
	book.GenreId = genre.Id

	err = a.bookrepo.Update(book)
	if err != nil {
		fmt.Printf("book update failed %v \n", req)
		return err
	}

	fmt.Printf("book updated %v \n", req)
	return nil
}

func (a *bookSvc) Get(id uuid.UUID) (*schema.Book, error) {
	book, err := a.bookrepo.GetById(id)
	if err != nil {
		fmt.Printf("get book failed, id - %v \n", id)
		return nil, err
	}

	fmt.Printf("retrieved book %v \n", book)
	return book, nil
}

func (a *bookSvc) List(page schema.ListBooks) ([]schema.Book, error) {

	if page.Limit > schema.DefaultLimit {
		page.Limit = schema.DefaultLimit
	}

	books, err := a.bookrepo.List(page.Page*page.Limit, page.Limit, page.GenreId, page.AuthorId)
	if err != nil {
		fmt.Println("failed to fetch books")
		return nil, err
	}

	fmt.Println("retrieved books", books)
	return books, nil
}

func (a *bookSvc) Delete(id uuid.UUID) error {
	err := a.bookrepo.Delete(id)
	if err != nil {
		fmt.Printf("book deletion failed, id - %v \n", id)
		return err
	}

	fmt.Printf("book deleted, id - %v \n", id)
	return nil
}

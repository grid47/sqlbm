package exec

import (
	"encoding/json"
	"fmt"
	"ysqlbm/schema"
	c "ysqlbm/test/client"

	"github.com/google/uuid"
)

func Book(client *c.Client, author, genre uuid.UUID) {

	// Create Book
	createBookReq := schema.CreateBook{
		Name:     "John Doe",
		Detail:   "Famous Book",
		AuthorId: author,
		GenreId:  genre,
	}
	var createBookResp schema.Response
	code, err := client.Post("/book/create", createBookReq, &createBookResp)
	if err != nil {
		fmt.Println("Error creating book:", err)
		return
	}
	fmt.Println("B-1. Book Created, code -", code)

	bytes, _ := json.Marshal(createBookResp.Data)
	book := schema.Book{}
	json.Unmarshal(bytes, &book)

	// Update Book
	updateBookReq := schema.UpdateBook{
		Id:       book.Id,
		Name:     "Updated Book",
		Detail:   "Updated Details",
		AuthorId: author,
		GenreId:  genre,
	}
	var updateBookResp schema.Response
	code, err = client.Post("/book/update", updateBookReq, &updateBookResp)
	if err != nil {
		fmt.Println("Error updating book:", err)
		return
	}
	fmt.Println("B-2. Book Updated, code -", code)

	var getBookResp schema.Response
	code, err = client.Get("/book/get",
		map[string]string{"id": book.Id.String()},
		&getBookResp)
	if err != nil || code != 200 {
		fmt.Println("Error getting book:", err, code)
		return
	}

	bytes, _ = json.Marshal(getBookResp.Data)
	newAuth := schema.Book{}
	json.Unmarshal(bytes, &newAuth)

	fmt.Printf("B-3. The Book %v, code %d \n", newAuth, code)

	listReq := schema.ListBooks{
		Page:     0,
		Limit:    4,
		GenreId:  genre.String(),
		AuthorId: author.String(),
	}
	var listBookResp schema.Response
	code, err = client.Post("/book/list", listReq, &listBookResp)
	if err != nil || code != 200 {
		fmt.Println("Error listing authors:", err, code)
		return
	}

	bytes, _ = json.Marshal(listBookResp.Data)
	books := []schema.Book{}
	json.Unmarshal(bytes, &books)

	fmt.Printf("B-4. List Books %v, code %d \n", books, code)

	var deleteBookResp schema.Response
	code, err = client.Get("/book/delete", map[string]string{"id": book.Id.String()}, &deleteBookResp)
	if err != nil {
		fmt.Println("Error deleting book:", err)
	}
	fmt.Printf("B-5. Book Deleted, code - %d \n", code)
}

package exec

import (
	"encoding/json"
	"fmt"
	"ysqlbm/schema"
	c "ysqlbm/test/client"

	"github.com/google/uuid"
)

func Genre(client *c.Client) *uuid.UUID {

	// Create Genre
	createGenreReq := schema.CreateGenre{
		Name:   "Genre",
		Detail: "Famous Genre",
	}
	var createGenreResp schema.Response
	code, err := client.Post("/genre/create", createGenreReq, &createGenreResp)
	if err != nil {
		fmt.Println("Error creating genre:", err)
		return nil
	}
	fmt.Println("G-1. Genre Created, code -", code)

	bytes, _ := json.Marshal(createGenreResp.Data)
	auth := schema.Genre{}
	json.Unmarshal(bytes, &auth)

	// Update Genre
	updateGenreReq := schema.UpdateGenre{
		Id:     auth.Id,
		Name:   "Updated Genre",
		Detail: "Updated Details",
	}
	var updateGenreResp schema.Response
	code, err = client.Post("/genre/update", updateGenreReq, &updateGenreResp)
	if err != nil {
		fmt.Println("Error updating genre:", err)
		return nil
	}
	fmt.Println("G-2. Genre Updated, code -", code)

	var getGenreResp schema.Response
	code, err = client.Get("/genre/get",
		map[string]string{"id": auth.Id.String()},
		&getGenreResp)
	if err != nil || code != 200 {
		fmt.Println("Error getting genre:", err, code)
		return nil
	}

	bytes, _ = json.Marshal(getGenreResp.Data)
	newAuth := schema.Genre{}
	json.Unmarshal(bytes, &newAuth)

	fmt.Printf("G-3. The Genre %v, code %d \n", newAuth, code)

	listReq := schema.Page{
		Page:  0,
		Limit: 4,
	}
	var listGenreResp schema.Response
	code, err = client.Post("/genre/list", listReq, &listGenreResp)
	if err != nil || code != 200 {
		fmt.Println("Error listing genres:", err, code)
		return nil
	}

	bytes, _ = json.Marshal(listGenreResp.Data)
	auths := []schema.Genre{}
	json.Unmarshal(bytes, &auths)

	fmt.Printf("G-4. List Genres %v, code %d \n", auths, code)

	return &auth.Id
}

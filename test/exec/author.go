package exec

import (
	"encoding/json"
	"fmt"
	"ysqlbm/schema"
	c "ysqlbm/test/client"

	"github.com/google/uuid"
)

func Author(client *c.Client) *uuid.UUID {

	// Create Author
	createAuthorReq := schema.CreateAuthor{
		Name:   "John Doe",
		Detail: "Famous Author",
	}
	var createAuthorResp schema.Response
	code, err := client.Post("/author/create", createAuthorReq, &createAuthorResp)
	if err != nil {
		fmt.Println("Error creating author:", err)
		return nil
	}
	fmt.Println("A-1. Author Created, code -", code)

	bytes, _ := json.Marshal(createAuthorResp.Data)
	auth := schema.Author{}
	json.Unmarshal(bytes, &auth)

	// Update Author
	updateAuthorReq := schema.UpdateAuthor{
		Id:     auth.Id,
		Name:   "Updated Author",
		Detail: "Updated Details",
	}
	var updateAuthorResp schema.Response
	code, err = client.Post("/author/update", updateAuthorReq, &updateAuthorResp)
	if err != nil {
		fmt.Println("Error updating author:", err)
		return nil
	}
	fmt.Println("A-2. Author Updated, code -", code)

	var getAuthorResp schema.Response
	code, err = client.Get("/author/get",
		map[string]string{"id": auth.Id.String()},
		&getAuthorResp)
	if err != nil || code != 200 {
		fmt.Println("Error getting author:", err, code)
		return nil
	}

	bytes, _ = json.Marshal(getAuthorResp.Data)
	newAuth := schema.Author{}
	json.Unmarshal(bytes, &newAuth)

	fmt.Printf("A-3. The Author %v, code %d \n", newAuth, code)

	listReq := schema.Page{
		Page:  0,
		Limit: 4,
	}
	var listAuthorResp schema.Response
	code, err = client.Post("/author/list", listReq, &listAuthorResp)
	if err != nil || code != 200 {
		fmt.Println("Error listing authors:", err, code)
		return nil
	}

	bytes, _ = json.Marshal(listAuthorResp.Data)
	auths := []schema.Author{}
	json.Unmarshal(bytes, &auths)

	fmt.Printf("A-4. List Authors %v, code %d \n", auths, code)
	return &auth.Id
}

func DelAuthor(id uuid.UUID, client *c.Client) {
	var deleteAuthorResp schema.Response
	code, err := client.Get("/author/delete", map[string]string{"id": id.String()}, &deleteAuthorResp)
	if err != nil {
		fmt.Println("Error deleting author:", err)
	}
	fmt.Printf("A-5. Author Deleted, code - %d \n", code)
}

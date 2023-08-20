package main

import (
	"ysqlbm/test/client"
	"ysqlbm/test/exec"
)

func main() {
	client := client.NewClient("http://localhost:9090")
	author := exec.Author(client)
	book := exec.Genre(client)
	exec.Book(client, *author, *book)
	exec.DelAuthor(*author, client)
}

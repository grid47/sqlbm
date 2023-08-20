package main

import (
	"fmt"
	"net/http"
	"ysqlbm/app/ctrl"
	"ysqlbm/app/repo"
	"ysqlbm/app/route"
	"ysqlbm/app/svc"
	"ysqlbm/db"
	"ysqlbm/schema"

	"github.com/gin-gonic/gin"
)

type server struct {
	cdb    *db.CDB
	router *gin.Engine
}

func main() {

	dsn := "host=localhost port=5432 user=myuser password=mypassword dbname=mydatabase sslmode=disable"

	cdb, err := db.NewCDB(dsn,
		&schema.Author{},
		&schema.Book{},
		&schema.Genre{})
	if err != nil {
		fmt.Println("db init failed", err)
		return
	}

	authorRepo := repo.NewAuthorDAL(cdb)
	authorSvc := svc.NewAuthorSvc(authorRepo)
	authorCtrl := ctrl.NewAuthorCtrl(authorSvc)

	genreRepo := repo.NewGenreDAL(cdb)
	genreSvc := svc.NewGenreSvc(genreRepo)
	genreCtrl := ctrl.NewGenreCtrl(genreSvc)

	bookRepo := repo.NewBookDAL(cdb)
	bookSvc := svc.NewBookSvc(bookRepo, genreRepo, authorRepo)
	bookCtrl := ctrl.NewBookCtrl(bookSvc)

	ctrls := ctrl.Controllers{
		AuthorCtrl: authorCtrl,
		BookCtrl:   bookCtrl,
		GenreCtrl:  genreCtrl,
	}

	srvr := server{
		cdb:    cdb,
		router: gin.New(),
	}

	srvr.router = route.Route(srvr.router, ctrls)
	handle := http.Handler(srvr.router)
	server := &http.Server{
		Addr:    ":9090",
		Handler: handle,
	}

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

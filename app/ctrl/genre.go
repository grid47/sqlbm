package ctrl

import (
	"net/http"
	"ysqlbm/app/svc"
	"ysqlbm/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GenreCtrl interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
}

var _ GenreCtrl = new(genreCtrl)

func NewGenreCtrl(genreSvc svc.GenreSvc) GenreCtrl {
	return &genreCtrl{
		genreSvc: genreSvc,
	}
}

type genreCtrl struct {
	genreSvc svc.GenreSvc
}

func (a *genreCtrl) Create(c *gin.Context) {
	var resp schema.Response

	req := schema.CreateGenre{}

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := a.genreSvc.Create(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: data}
	c.JSON(http.StatusOK, resp)
}

func (a *genreCtrl) Update(c *gin.Context) {
	var resp schema.Response

	req := schema.UpdateGenre{}

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := a.genreSvc.Update(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: "Genre updated successfully"}
	c.JSON(http.StatusOK, resp)
}

func (a *genreCtrl) Get(c *gin.Context) {
	var resp schema.Response

	idStr := c.Query("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		resp = schema.Response{Error: "Invalid UUID format"}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	author, err := a.genreSvc.Get(id)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp = schema.Response{Data: author}
	c.JSON(http.StatusOK, resp)
}

func (a *genreCtrl) List(c *gin.Context) {
	var resp schema.Response
	var req schema.Page

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	authors, err := a.genreSvc.List(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: authors}
	c.JSON(http.StatusOK, resp)
}

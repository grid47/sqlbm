package ctrl

import (
	"net/http"
	"ysqlbm/app/svc"
	"ysqlbm/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookCtrl interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
}

var _ BookCtrl = new(bookCtrl)

func NewBookCtrl(bookSvc svc.BookSvc) BookCtrl {
	return &bookCtrl{
		bookSvc: bookSvc,
	}
}

type bookCtrl struct {
	bookSvc svc.BookSvc
}

func (a *bookCtrl) Create(c *gin.Context) {
	var resp schema.Response

	req := schema.CreateBook{}

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := a.bookSvc.Create(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: data}
	c.JSON(http.StatusOK, resp)
}

func (a *bookCtrl) Update(c *gin.Context) {
	var resp schema.Response

	req := schema.UpdateBook{}

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := a.bookSvc.Update(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: "Book updated successfully"}
	c.JSON(http.StatusOK, resp)
}

func (a *bookCtrl) Get(c *gin.Context) {
	var resp schema.Response

	idStr := c.Query("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		resp = schema.Response{Error: "Invalid UUID format"}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	book, err := a.bookSvc.Get(id)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp = schema.Response{Data: book}
	c.JSON(http.StatusOK, resp)
}

func (a *bookCtrl) List(c *gin.Context) {
	var resp schema.Response
	var req schema.ListBooks

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	books, err := a.bookSvc.List(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: books}
	c.JSON(http.StatusOK, resp)
}

func (a *bookCtrl) Delete(c *gin.Context) {
	var resp schema.Response

	idStr := c.Query("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		resp = schema.Response{Error: "Invalid UUID format"}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err = a.bookSvc.Delete(id)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: "Book deleted successfully"}
	c.JSON(http.StatusOK, resp)
}

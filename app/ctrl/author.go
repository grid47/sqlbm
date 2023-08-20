package ctrl

import (
	"fmt"
	"net/http"
	"ysqlbm/app/svc"
	"ysqlbm/schema"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthorCtrl interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Get(c *gin.Context)
	List(c *gin.Context)
	Delete(c *gin.Context)
}

var _ AuthorCtrl = new(authorCtrl)

func NewAuthorCtrl(authorSvc svc.AuthorSvc) AuthorCtrl {
	return &authorCtrl{
		authorSvc: authorSvc,
	}
}

type authorCtrl struct {
	authorSvc svc.AuthorSvc
}

func (a *authorCtrl) Create(c *gin.Context) {
	var resp schema.Response

	req := schema.CreateAuthor{}

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := a.authorSvc.Create(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: data}
	c.JSON(http.StatusOK, resp)
}

func (a *authorCtrl) Update(c *gin.Context) {
	var resp schema.Response

	req := schema.UpdateAuthor{}

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err := a.authorSvc.Update(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: "Author updated successfully"}
	c.JSON(http.StatusOK, resp)
}

func (a *authorCtrl) Get(c *gin.Context) {
	var resp schema.Response

	idStr := c.Query("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		resp = schema.Response{Error: "Invalid UUID format"}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	fmt.Printf("get author %v", id)

	author, err := a.authorSvc.Get(id)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp = schema.Response{Data: author}
	c.JSON(http.StatusOK, resp)
}

func (a *authorCtrl) List(c *gin.Context) {
	var resp schema.Response
	var req schema.Page

	fmt.Printf("list authors --")

	if err := c.BindJSON(&req); err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	fmt.Printf("list authors %v", req)

	authors, err := a.authorSvc.List(req)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: authors}
	c.JSON(http.StatusOK, resp)
}

func (a *authorCtrl) Delete(c *gin.Context) {
	var resp schema.Response

	idStr := c.Query("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		resp = schema.Response{Error: "Invalid UUID format"}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	err = a.authorSvc.Delete(id)
	if err != nil {
		resp = schema.Response{Error: err.Error()}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp = schema.Response{Data: "Author deleted successfully"}
	c.JSON(http.StatusOK, resp)
}

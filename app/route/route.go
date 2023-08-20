package route

import (
	"ysqlbm/app/ctrl"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine, env ctrl.Controllers) *gin.Engine {

	router.POST("/author/create", env.AuthorCtrl.Create)
	router.POST("/author/update", env.AuthorCtrl.Update)
	router.GET("/author/get", env.AuthorCtrl.Get)
	router.POST("/author/list", env.AuthorCtrl.List)
	router.GET("/author/delete", env.AuthorCtrl.Delete)

	router.POST("/genre/create", env.GenreCtrl.Create)
	router.POST("/genre/update", env.GenreCtrl.Update)
	router.GET("/genre/get", env.GenreCtrl.Get)
	router.POST("/genre/list", env.GenreCtrl.List)

	router.POST("/book/create", env.BookCtrl.Create)
	router.POST("/book/update", env.BookCtrl.Update)
	router.GET("/book/get", env.BookCtrl.Get)
	router.POST("/book/list", env.BookCtrl.List)
	router.GET("/book/delete", env.BookCtrl.Delete)

	return router
}

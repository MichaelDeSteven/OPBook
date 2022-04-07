package router

import (
	"github.com/MichaelDeSteven/OPBook/server/controller"
	"github.com/MichaelDeSteven/OPBook/server/middleware"
	"github.com/MichaelDeSteven/rum"
)

type BookRouter struct{}

func (s *BookRouter) InitBookRouter(r *rum.RouterGroup) {
	bookRouter := r.Group("book")
	bookRouter.Use(middleware.Recovery(true)).Use(middleware.DefaultLogger()).Use(middleware.JWTAuth())
	bookRouter.POST("/create", controller.Create)
	bookRouter.POST("/index", controller.Index)
	bookRouter.POST("/upload", controller.UploadProject)
	bookRouter.POST("/release/:bookId", controller.Release)
	bookRouter.GET("/introduct/:identify", controller.Introduct)
	bookRouter.POST("/collect/:bookId", controller.Star)
	bookRouter.GET("/collect/stat/:bookId", controller.IsStar)
	bookRouter.POST("/score/add", controller.AddScore)
	bookRouter.POST("/score/get", controller.GetScore)
}

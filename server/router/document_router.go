package router

import (
	"github.com/MichaelDeSteven/OPBook/server/controller"
	"github.com/MichaelDeSteven/OPBook/server/middleware"
	"github.com/MichaelDeSteven/rum"
)

type DocumentRouter struct{}

func (s *DocumentRouter) InitDocumentRouter(r *rum.RouterGroup) {
	docController := controller.NewDocumentController()
	documentRouter := r.Group("document")
	documentRouter.Use(middleware.Recovery(true)).Use(middleware.DefaultLogger()).Use(middleware.JWTAuth())
	documentRouter.GET("/index/:identify", docController.Index)
}

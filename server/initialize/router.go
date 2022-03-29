package initialize

import (
	"github.com/MichaelDeSteven/OPBook/server/router"
	"github.com/MichaelDeSteven/rum"
)

func initRouters() *rum.Engine {
	r := rum.Default()
	PublicGroup := r.Group("")
	base := router.BaseRouter{}
	user := router.UserRouter{}
	book := router.BookRouter{}
	doc := router.DocumentRouter{}
	base.InitBaseRouter(PublicGroup)
	user.InitUserRouter(PublicGroup)
	book.InitBookRouter(PublicGroup)
	doc.InitDocumentRouter(PublicGroup)
	return r
}

package initialize

import (
	"github.com/MichaelDeSteven/OPBook/server/router"
	"github.com/MichaelDeSteven/rum"
)

func Routers() *rum.Engine {
	r := rum.Default()
	PublicGroup := r.Group("")
	base := router.BaseRouter{}
	user := router.UserRouter{}
	book := router.BookRouter{}
	base.InitBaseRouter(PublicGroup)
	user.InitUserRouter(PublicGroup)
	book.InitBookRouter(PublicGroup)
	return r
}

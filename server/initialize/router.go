package initialize

import (
	"github.com/MichaelDeSteven/OPBook/server/router"
	"github.com/MichaelDeSteven/rum"
)

func Routers() *rum.Engine {
	r := rum.Default()
	PublicGroup := r.Group("")
	user := router.UserRouter{}
	book := router.BookRouter{}
	user.InitUserRouter(PublicGroup)
	book.InitBookRouter(PublicGroup)
	return r
}

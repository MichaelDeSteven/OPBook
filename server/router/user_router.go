package router

import (
	"github.com/MichaelDeSteven/OPBook/server/controller"
	"github.com/MichaelDeSteven/OPBook/server/middleware"
	"github.com/MichaelDeSteven/rum"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(r *rum.RouterGroup) {
	userRouter := r.Group("user")
	userRouter.Use(middleware.Recovery(true)).Use(middleware.DefaultLogger()).Use(middleware.JWTAuth())
	userRouter.POST("/update", controller.UpdateUserInfo)
	userRouter.POST("/setting/upload", controller.UploadAvatar)
	userRouter.POST("/password", controller.Password)
}

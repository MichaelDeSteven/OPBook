package router

import (
	"github.com/MichaelDeSteven/OPBook/server/controller"
	"github.com/MichaelDeSteven/OPBook/server/middleware"
	"github.com/MichaelDeSteven/rum"
)

type SocialRouter struct{}

func (s *SocialRouter) InitSocialRouter(r *rum.RouterGroup) {
	socialRouter := r.Group("social")
	socialRouter.Use(middleware.Recovery(true)).Use(middleware.DefaultLogger()).Use(middleware.JWTAuth())
	follow := controller.FollowController{}
	comment := controller.CommentController{}
	chat := controller.ChatController{}
	socialRouter.POST("/fan/follow", follow.Follow)
	socialRouter.POST("/fan/stat", follow.FollowStatus)
	socialRouter.POST("/fan/getFollowees", follow.GetFollowees)
	socialRouter.POST("/fan/getFollowers", follow.GetFollowers)
	socialRouter.POST("/comment/add", comment.CommentOrReply)
	socialRouter.GET("/comment/get/:bookId", comment.DisplayComment)
	socialRouter.POST("/chat/add", chat.SendMessage)
	socialRouter.GET("/chat/getUserList", chat.GetConversationUserList)
	socialRouter.POST("/chat/getConversation", chat.GetConversation)
	socialRouter.POST("/like/:commentId", comment.Like)
}

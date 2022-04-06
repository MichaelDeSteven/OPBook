package controller

import (
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/rum"
)

type ChatController struct {
}

func (this *ChatController) SendMessage(c *rum.Context) {
	message := model.NewMessage()
	c.Bind(message)
	socialService.SendMessage(message)
}

func (this *ChatController) GetConversationUserList(c *rum.Context) {
	uid, _ := c.Get("uid")
	if uid == nil {
		response.FailWithMessage("用户未登录", c)
		return
	}
	response.OkWithData(socialService.GetConversationUserList(uid.(int)), c)
}

func (this *ChatController) GetConversation(c *rum.Context) {
	message := model.NewMessage()
	c.Bind(message)
	response.OkWithData(socialService.GetConversation(message), c)
}

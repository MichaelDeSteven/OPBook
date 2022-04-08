package controller

import (
	"strconv"

	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/rum"
)

type CommentController struct {
}

func (this *CommentController) CommentOrReply(c *rum.Context) {
	comment := model.NewComment()
	c.Bind(comment)
	if comment.BookId <= 0 {
		response.FailWithMessage("bookId为空", c)
		return
	}
	socialService.CommentOrReply(*comment)
}

func (this *CommentController) DisplayComment(c *rum.Context) {
	bid, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		response.FailWithMessage("bookId应该为数字", c)
		return
	}
	login := false
	uid, _ := c.Get("uid")
	if uid.(int) > 0 {
		login = true
	}
	response.OkWithData(socialService.DisplayComment(bid, uid.(int), login), c)
}

func (this *CommentController) Like(c *rum.Context) {
	uid, _ := c.Get("uid")
	if uid.(int) < 1 {
		response.FailWithMessage("用户未登录", c)
		return
	}
	commentId, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		response.FailWithMessage("commentId应该为数字", c)
		return
	}
	response.OkWithData(socialService.Like(commentId, uid.(int)), c)
}

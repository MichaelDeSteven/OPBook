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
	response.OkWithData(socialService.DisplayComment(bid), c)
}

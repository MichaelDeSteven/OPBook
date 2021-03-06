package controller

import (
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/MichaelDeSteven/rum"
)

type DocumentController struct {
}

func NewDocumentController() *DocumentController {
	return &DocumentController{}
}

// 文档目录
func (this *DocumentController) Index(c *rum.Context) {
	identify := c.Param("identify")
	if utils.Empty(identify) {
		response.FailWithMessage("identify为空", c)
		return
	}

	bookId := model.NewBook().FindIdByIdentify(identify)
	docs := model.NewDocument().GetMenuTop(bookId)
	response.OkWithData(docs, c)
}

// 阅读文档
func (this *DocumentController) Read(c *rum.Context) {
	bookIdentify, docIdentify := c.Param("book_identify"), c.Param("doc_identify")
	if utils.Empty(bookIdentify) || utils.Empty(docIdentify) {
		response.FailWithMessage("identify为空", c)
		return
	}
	data, err := bookService.Read(bookIdentify, docIdentify)
	if err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithData(data, c)
}

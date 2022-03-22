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
	book := model.NewBook().FindColsByIdentify(bookIdentify, "id", "author", "author_url", "name")
	doc := model.NewDocument().Get(book.Id, docIdentify)
	if doc == nil || doc.Id == 0 {
		response.FailWithMessage("目标文档不存在", c)
		return
	}
	data := make(map[string]interface{})
	data["doc"] = doc
	tree, err := model.NewDocument().CreateDocumentTreeForHtml(doc.BookId, doc.Id)
	if err != nil {
		response.FailWithMessage("获取书籍目录失败", c)
		return
	}
	data["menu"] = tree
	data["book"] = book
	response.OkWithData(data, c)
}

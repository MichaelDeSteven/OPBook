package controller

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/model/response"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/MichaelDeSteven/rum"
)

// Create 创建书籍.
func Create(c *rum.Context) {
	book := model.NewBook()
	c.Bind(book)
	book.Name = strings.TrimSpace(book.Name)
	book.Identify = strings.TrimSpace(book.Identify)
	book.Description = strings.TrimSpace(book.Description)
	book.Author = strings.TrimSpace(book.Author)
	book.Author = strings.TrimSpace(book.Author)
	book.AuthorURL = strings.TrimSpace(book.AuthorURL)

	if utils.Empty(book.Name) {
		response.FailWithMessage("书籍名称不能为空", c)
		return
	}

	if utils.Empty(book.Identify) {
		response.FailWithMessage("书籍标识不能为空", c)
		return
	}

	ok, err1 := regexp.MatchString(`^[a-zA-Z0-9_\-\.]*$`, book.Identify)
	if !ok || err1 != nil {
		response.FailWithMessage("书籍标识只能包含字母、数字，以及“-”、“.”和“_”符号，且不能是纯数字", c)
		return
	}

	if num, _ := strconv.Atoi(book.Identify); strconv.Itoa(num) == book.Identify {
		response.FailWithMessage("书籍标识不能是纯数字", c)
		return
	}

	if strings.Count(book.Identify, "") > 50 {
		response.FailWithMessage("书籍标识不能超过50字", c)
		return
	}

	if strings.Count(book.Description, "") > 500 {
		response.FailWithMessage("书籍描述不能大于500字", c)
		return
	}
	if err := bookService.Create(book); err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

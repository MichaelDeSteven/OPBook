package controller

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/MichaelDeSteven/OPBook/server/global"
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

//上传书籍
func UploadProject(c *rum.Context) {
	//处理步骤
	//1、接受上传上来的zip文件，并存放到store/temp目录下
	//2、解压zip到当前目录，然后移除非图片文件
	//3、将文件夹移动到uploads目录下
	form := &model.BookForm{}
	c.Bind(form)
	uid, _ := c.Get("uid")
	if !model.NewBook().HasProjectAccess(form.Identify, uid.(int)) {
		response.FailWithMessage("无操作权限", c)
		return
	}

	// book, _ := model.NewBookResult().FindByIdentify(form.Identify, uid.(int))
	// if book.BookId == 0 {
	// 	response.FailWithMessage("书籍不存在", c)
	// 	return
	// }

	// f, h, err := this.GetFile("zipfile")
	// if err != nil {
	// 	this.JsonResult(1, err.Error())
	// }
	// defer f.Close()
	// if strings.ToLower(filepath.Ext(h.Filename)) != ".zip" && strings.ToLower(filepath.Ext(h.Filename)) != ".epub" {
	// 	this.JsonResult(1, "请上传指定格式文件")
	// }
	// tmpFile := "store/" + identify + ".zip" //保存的文件名
	// if err := this.SaveToFile("zipfile", tmpFile); err == nil {
	// 	go unzipToData(book.BookId, identify, tmpFile, h.Filename)
	// } else {
	// 	beego.Error(err.Error())
	// }
	// this.JsonResult(0, "上传成功")
	response.OkWithMessage("上传成功", c)
}

func Index(c *rum.Context) {
	req := model.BookPage{}
	c.Bind(&req)
	global.LOG.Sugar().Infof("%+v\n", req)
	books, totalCount, err := model.NewBook().FindToPager(req.PageIndex, req.PageSize, 0)
	if err != nil {
		global.LOG.Sugar().Errorf("查询失败: %+v\n", err)
		response.FailWithMessage("查询失败", c)
		return
	}
	data := make(map[string]interface{})
	data["books"] = books
	data["totalCount"] = totalCount
	response.OkWithData(data, c)
}

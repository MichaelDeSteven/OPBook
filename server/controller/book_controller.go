package controller

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

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

	book := model.NewBook().FindByIdentify(form.Identify)
	if book.Id == 0 {
		response.FailWithMessage("书籍不存在", c)
		return
	}

	h := form.Zip
	if strings.ToLower(filepath.Ext(h.Filename)) != ".zip" {
		response.FailWithMessage("请上传指定格式文件", c)
		return
	}
	filePath := filepath.Join("./", global.CONFIG.Local.Path, "tmp", global.CONFIG.Local.Book, time.Now().Format("2006/01"), h.Filename)
	path := filepath.Dir(filePath)
	os.MkdirAll(path, os.ModePerm)
	if err := c.SaveUploadedFile(h, filePath); err != nil {
		global.LOG.Sugar().Errorf("保存文件失败: %+v\n", err)
		response.FailWithMessage("保存文件失败", c)
		return
	} else {
		go bookService.UnzipToData(book.Id, form.Identify, filePath, h.Filename)
	}

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

//发布书籍.
func Release(c *rum.Context) {
	bookId, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		response.FailWithMessage("bookId应该为数字", c)
		return
	}
	model.NewDocument().ReleaseContent(bookId)
}

// 书籍概要
func Introduct(c *rum.Context) {
	identify := c.Param("identify")
	if utils.Empty(identify) {
		response.FailWithMessage("identify为空", c)
		return
	}
	book := bookService.GetBookIntroduct(identify)
	response.OkWithData(book, c)
}

// 收藏或者取消收藏书籍
func Star(c *rum.Context) {
	uid, _ := c.Get("uid")
	if uid == nil {
		response.FailWithMessage("用户未登录", c)
		return
	}
	bid, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		response.FailWithMessage("bookId应该为数字", c)
		return
	}
	if cancel := bookService.Star(uid.(int), bid); cancel {
		response.OkWithMessage("已取消收藏", c)
	} else {
		response.OkWithMessage("已收藏", c)
	}
}

// 用户收藏书籍的状态
func IsStar(c *rum.Context) {
	uid, _ := c.Get("uid")
	model := &model.Star{
		IsDeleted: 1,
	}
	// 游客状态则未被收藏
	if uid == nil {
		response.OkWithData(model, c)
		return
	}
	bid, err := strconv.Atoi(c.Param("bookId"))
	if err != nil {
		response.FailWithMessage("bookId应该为数字", c)
		return
	}
	if bookService.IsStar(uid.(int), bid) {
		model.IsDeleted = 0
		response.OkWithData(model, c)
	} else {
		response.OkWithData(model, c)
	}
}

// 用户书籍收藏列表
func UserCollection(c *rum.Context) {
	page := model.NewUserCollectPage()
	c.Bind(&page)
	global.LOG.Sugar().Infof("%+v\n", page)
	books, totalCount, err := bookService.UserCollection(page)
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

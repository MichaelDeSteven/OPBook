package service

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/model"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/MichaelDeSteven/OPBook/server/utils/html2md"
	"github.com/MichaelDeSteven/OPBook/server/utils/upload"
	"github.com/PuerkitoBio/goquery"
	"github.com/TruthHun/gotil/filetil"
	"github.com/TruthHun/gotil/mdtil"
	"github.com/TruthHun/gotil/ziptil"
	"github.com/go-redis/redis/v8"
	"github.com/russross/blackfriday"
)

type BookService struct{}

func (bookservice *BookService) Create(b *model.Book) error {
	if b.CheckIdentifyIsExist(b.Identify) {
		return errors.New("书籍唯一标识已存在")
	}
	defaultTime, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	b.LastClickGenerate = defaultTime
	b.GenerateTime, _ = time.Parse("2006-01-02 15:04:05", "2000-01-02 15:04:05") // 默认生成文档的时间
	b.ReleaseTime = defaultTime

	b.Version = time.Now().Unix()
	return b.Add(b)
}

func (bookservice *BookService) UnzipToData(bookId int, identify, zipFile, originFilename string) {

	oss := upload.NewOss()
	imgMap := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".bmp": true, ".svg": true, ".webp": true}
	projectRoot := "" //书籍根目录
	// 解压目录
	unzipPath := global.CONFIG.Local.Path + "/store/" + identify

	// 如果存在相同目录，则率先移除
	if err := os.RemoveAll(unzipPath); err != nil {
		global.LOG.Sugar().Errorf("删除目录失败: %+v\n", err)
		return
	}
	os.MkdirAll(unzipPath, os.ModePerm)
	defer func() {
		os.Remove(zipFile)      // 最后删除上传的临时文件
		os.RemoveAll(unzipPath) // 删除解压后的文件夹
	}()

	if err := ziptil.Unzip(zipFile, unzipPath); err != nil {
		global.LOG.Sugar().Errorf("解压失败: %+v\n err:%+v\n", zipFile, err)
		return
	}

	// 读取文件，把图片文档录入oss
	if files, err := filetil.ScanFiles(unzipPath); err == nil {
		projectRoot = getProjectRoot(files)
		replaceToAbs(projectRoot, identify)
		modelStore := new(model.DocumentStore)
		for _, file := range files {
			if !file.IsDir {
				ext := strings.ToLower(filepath.Ext(file.Path))
				if ok, _ := imgMap[ext]; ok {
					if err := oss.UploadFileByPrefix(file.Path, "projects/"+identify+strings.TrimPrefix(file.Path, projectRoot)); err != nil {
						global.LOG.Sugar().Errorf("err:%+v\n", err)
					}
				} else if ext == ".md" || ext == ".markdown" || ext == ".html" {
					// markdown文档，提取文档内容，录入数据库
					doc := model.NewDocument()
					var mdcont string
					var htmlStr string
					if b, err := ioutil.ReadFile(file.Path); err == nil {
						if ext == ".md" || ext == ".markdown" {
							mdcont = strings.TrimSpace(string(b))
							htmlStr = mdtil.Md2html(mdcont)
						} else {
							htmlStr = string(b)
							mdcont = html2md.Convert(htmlStr)
						}
						// if !strings.HasPrefix(mdcont, "[TOC]") {
						// 	mdcont = "[TOC]\r\n\r\n" + mdcont
						// }
						doc.Name = utils.ParseTitleFromMdHtml(htmlStr)
						doc.BookId = bookId
						doc.Identify = strings.Replace(strings.Trim(strings.TrimPrefix(file.Path, ""), "/"), "/", "-", -1)
						doc.Identify = strings.Replace(doc.Identify, ")", "", -1)
						doc.UserId = 0
						doc.OrderSort = 1
						if strings.HasSuffix(strings.ToLower(file.Name), "summary.md") {
							doc.OrderSort = 0
						}
						if strings.HasSuffix(strings.ToLower(file.Name), "summary.html") {
							mdcont += "<bookstack-summary></bookstack-summary>"
							// 生成带$的文档标识，阅读BaseController.go代码可知，
							// 要使用summary.md的排序功能，必须在链接中带上符号$
							mdcont = strings.Replace(mdcont, "](", "]($", -1)
							// 去掉可能存在的url编码的右括号，否则在url译码后会与markdown语法混淆
							mdcont = strings.Replace(mdcont, "%29", "", -1)
							mdcont, _ = url.QueryUnescape(mdcont)
							doc.OrderSort = 0
							doc.Identify = "summary.md"
						}
						global.LOG.Sugar().Infof("doc:%+v\n", doc)
						if docId, err := doc.InsertOrUpdate(); err == nil {
							if err := modelStore.InsertOrUpdate(&model.DocumentStore{
								DocumentId: docId,
								Markdown:   mdcont,
								Content:    mdtil.Md2html(mdcont)}, "markdown", "content"); err != nil {
								global.LOG.Sugar().Errorf("err:%+v\n", err)
							}
						} else {
							global.LOG.Sugar().Errorf("err:%+v\n", err)
						}
					} else {
						global.LOG.Sugar().Errorf("读取文档失败：: %+v\n err:%+v\n", file.Path, err)
					}
				}
			}
		}
	}
}

// 获取书籍的根目录
func getProjectRoot(fl []filetil.FileList) (root string) {
	i := 1000
	for _, f := range fl {
		if !f.IsDir {
			if cnt := strings.Count(f.Path, "/"); cnt < i {
				root = filepath.Dir(f.Path)
				i = cnt
			}
		}
	}
	return
}

// 查找并替换markdown文件中的路径，把图片链接替换成url的相对路径，把文档间的链接替换成【$+文档标识链接】
func replaceToAbs(projectRoot string, identify string) {
	imgBaseUrl := global.CONFIG.TencentCOS.BaseURL
	files, _ := filetil.ScanFiles(projectRoot)
	for _, file := range files {
		if ext := strings.ToLower(filepath.Ext(file.Path)); ext == ".md" || ext == ".markdown" {
			// mdb ==> markdown byte
			mdb, _ := ioutil.ReadFile(file.Path)
			mdCont := string(mdb)
			basePath := filepath.Dir(file.Path)
			basePath = strings.Trim(strings.Replace(basePath, "\\", "/", -1), "/")
			basePathSlice := strings.Split(basePath, "/")
			l := len(basePathSlice)
			b, _ := ioutil.ReadFile(file.Path)
			output := blackfriday.Run(b)
			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(output)))

			// 图片链接处理
			doc.Find("img").Each(func(i int, selection *goquery.Selection) {
				// 非http://、// 和 https:// 开头的图片地址，即是相对地址
				src, ok := selection.Attr("src")
				lowerSrc := strings.ToLower(src)
				if ok &&
					!strings.HasPrefix(lowerSrc, "http://") &&
					!strings.HasPrefix(lowerSrc, "https://") {
					newSrc := src //默认为旧地址
					if strings.HasPrefix(lowerSrc, "//") {
						newSrc = "https:" + newSrc
					} else {
						if cnt := strings.Count(src, "../"); cnt < l { // 替换以或者"../"开头的路径
							newSrc = strings.Join(basePathSlice[0:l-cnt], "/") + "/" + strings.TrimLeft(src, "./")
						}
						newSrc = imgBaseUrl + "/" + strings.TrimLeft(strings.TrimPrefix(strings.TrimLeft(newSrc, "./"), projectRoot), "/")
					}
					mdCont = strings.Replace(mdCont, src, newSrc, -1)
				}
			})

			// a标签链接处理。要注意判断有锚点的情况
			doc.Find("a").Each(func(i int, selection *goquery.Selection) {
				href, ok := selection.Attr("href")
				lowerHref := strings.TrimSpace(strings.ToLower(href))
				// 链接存在，且不以 // 、 http、https、mailto 开头
				if ok &&
					!strings.HasPrefix(lowerHref, "//") &&
					!strings.HasPrefix(lowerHref, "http://") &&
					!strings.HasPrefix(lowerHref, "https://") &&
					!strings.HasPrefix(lowerHref, "mailto:") &&
					!strings.HasPrefix(lowerHref, "#") {
					newHref := href //默认
					if cnt := strings.Count(href, "../"); cnt < l {
						newHref = strings.Join(basePathSlice[0:l-cnt], "/") + "/" + strings.TrimLeft(href, "./")
					}
					newHref = strings.TrimPrefix(strings.Trim(newHref, "/"), projectRoot)
					if !strings.HasPrefix(href, "$") { // 原链接不包含$符开头，否则表示已经替换过了。
						newHref = "$" + strings.Replace(strings.Trim(newHref, "/"), "/", "-", -1)
						slice := strings.Split(newHref, "$")
						if ll := len(slice); ll > 0 {
							newHref = "$" + slice[ll-1]
						}
						mdCont = strings.Replace(mdCont, "]("+href, "]("+newHref, -1)
					}
				}
			})
			ioutil.WriteFile(file.Path, []byte(mdCont), os.ModePerm)
		}
	}
}

// 获取书籍概括
func (bookservice *BookService) GetBookIntroduct(identify string) *model.Book {
	book, err := model.NewBook().Get(identify)
	if err != nil {
		global.LOG.Sugar().Error(err)
		return nil
	}
	increaseBookView(book.Id)
	book.CollectCount = getCount(utils.BOOK_STAR_COUNT_PREFIX + fmt.Sprint(book.Id))
	book.ViewCount = getCount(utils.BOOK_VIEW_COUNT_PREFIX + fmt.Sprint(book.Id))
	return book
}

// 在线阅读
func (bookservice *BookService) Read(bookIdentify, docIdentify string) (map[string]interface{}, error) {
	book := model.NewBook().FindColsByIdentify(bookIdentify, "id", "author", "author_url", "name")
	doc := model.NewDocument().Get(book.Id, docIdentify)
	if doc == nil || doc.Id == 0 {
		global.LOG.Sugar().Errorf("目标文档不存在")
		return nil, errors.New("目标文档不存在")
	}
	data := make(map[string]interface{}, 3)
	doc.ViewCount = getCount(utils.DOC_VIEW_COUNT_PREFIX + fmt.Sprint(doc.Id))
	data["doc"] = doc
	tree, err := model.NewDocument().CreateDocumentTreeForHtml(doc.BookId, doc.Id)
	if err != nil {
		global.LOG.Sugar().Errorf("获取书籍目录失败：%+v", err)
		return nil, err
	}
	data["menu"] = tree
	data["book"] = book
	increaseDocView(doc.Id)
	return data, nil
}

func increaseBookView(bookId int) {
	key := utils.BOOK_VIEW_COUNT_PREFIX + fmt.Sprint(bookId)
	val, err := global.REDIS.Get(context.Background(), key).Int()
	if err == nil {
		global.REDIS.Set(context.Background(), key, val+1, 0)
	} else if err == redis.Nil {
		global.REDIS.Set(context.Background(), key, model.NewBook().GetBookView(bookId)+1, 0)
	} else {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}
}

func increaseDocView(docId int) {
	key := utils.DOC_VIEW_COUNT_PREFIX + fmt.Sprint(docId)
	val, err := global.REDIS.Get(context.Background(), key).Int()
	if err == nil {
		global.REDIS.Set(context.Background(), key, val+1, 0)
	} else if err == redis.Nil {
		global.REDIS.Set(context.Background(), key, model.NewDocument().GetDocumentView(docId)+1, 0)
	} else {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}
}

// 获取Redis缓存key的value
func getCount(key string) int {
	val, _ := global.REDIS.Get(context.Background(), key).Int()
	return val
}

func UpdateBooksView() {
	book := model.NewBook()
	set, _ := global.REDIS.Keys(context.Background(), utils.BOOK_VIEW_COUNT_PREFIX+"*").Result()
	for _, val := range set {
		ss := strings.Split(val, "_")
		bookId, _ := strconv.Atoi(ss[len(ss)-1])
		viewCount, _ := global.REDIS.Get(context.Background(), val).Int()
		book.SetBookView(bookId, viewCount)
	}
}

func UpdateDocsView() {
	doc := model.NewDocument()
	set, _ := global.REDIS.Keys(context.Background(), utils.DOC_VIEW_COUNT_PREFIX+"*").Result()
	for _, val := range set {
		ss := strings.Split(val, "_")
		docId, _ := strconv.Atoi(ss[len(ss)-1])
		viewCount, _ := global.REDIS.Get(context.Background(), val).Int()
		doc.SetDocumentView(docId, viewCount)
	}
}

// 收藏或者取消收藏
func (bookservice *BookService) Star(userId, bookId int) (cancel bool) {
	var model model.Star
	star := model.GetStarbyUserIdAndBookId(userId, bookId)
	// 取消收藏
	if star.Id == 0 {
		star.Star(userId, bookId, false)
		setIncreAndDecreCollectCount(bookId, 1)
	} else {
		if star.IsDeleted == 0 {
			star.IsDeleted = 1
			star.Star(userId, bookId, true)
			cancel = true
			setIncreAndDecreCollectCount(bookId, -1)
		} else {
			star.IsDeleted = 0
			star.Star(userId, bookId, true)
			setIncreAndDecreCollectCount(bookId, 1)
		}
	}
	return
}

func (bookservice *BookService) IsStar(userId, bookId int) bool {
	var model model.Star
	return model.IsStar(userId, bookId)
}

func setIncreAndDecreCollectCount(bookId, amount int) {
	key := utils.BOOK_STAR_COUNT_PREFIX + fmt.Sprint(bookId)
	val, err := global.REDIS.Get(context.Background(), key).Int()
	if err == nil {
		global.REDIS.Set(context.Background(), key, val+amount, 0)
	} else if err == redis.Nil {
		global.REDIS.Set(context.Background(), key, model.NewBook().GetCollectCount(bookId)+amount, 0)
	} else {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}
}

func UpdateBooksStar() {
	book := model.NewBook()
	set, _ := global.REDIS.Keys(context.Background(), utils.BOOK_STAR_COUNT_PREFIX+"*").Result()
	for _, val := range set {
		ss := strings.Split(val, "_")
		bookId, _ := strconv.Atoi(ss[len(ss)-1])
		collectCount, _ := global.REDIS.Get(context.Background(), val).Int()
		global.LOG.Sugar().Infof("%+v %v\n", bookId, collectCount)
		book.SetCollectCount(bookId, collectCount)
	}
}

func (bookservice *BookService) UserCollection(page *model.UserCollectPage) (books []*model.Book, totalCount int64, err error) {
	if page.PageIndex < 1 {
		page.PageIndex = 1
	}
	star, book := model.NewStar(), model.NewBook()
	totalCount = star.GetUserCollectCount(page.UserId)
	stars := star.GetUserCollectList(page.UserId, page.PageIndex, page.PageSize)
	ids := make([]int, len(stars))
	for i, star := range stars {
		ids[i] = star.BookId
	}
	books, _ = book.GetBooksById(ids)
	return
}

func (bookservice *BookService) AddScore(score *model.Score) {
	if score.GetScore(score.BookId, score.UserId).Id >= 1 {
		global.LOG.Sugar().Errorf("您已给当前书籍打过分了")
		return
	}
	score.AddScore()
	book := model.NewBook().GetScore(score.BookId)
	book.ScoreCount++
	book.Score = (book.Score*(book.ScoreCount-1) + score.Score) / book.ScoreCount
	book.SetScore(book)
}

func (bookservice *BookService) GetScore(score *model.Score) *model.Score {
	return score.GetScore(score.BookId, score.UserId)
}

func (bookservice *BookService) SetBook(b *model.Book) {
	b.UpdateBook(map[string]interface{}{"name": b.Name, "author": b.Author, "author_url": b.AuthorURL, "lang": b.Lang, "label": b.Label, "description": b.Description})
}

func (bookservice *BookService) UploadCover(b *model.Book) {
	b.UpdateBook(map[string]interface{}{"cover": b.Cover})
}

func (bookservice *BookService) GetBookId(bookId int) *model.Book {
	books, err := model.NewBook().GetBooksById([]int{bookId})
	if err != nil {
		global.LOG.Sugar().Error(err)
		return nil
	}
	if len(books) < 1 {
		global.LOG.Sugar().Errorf("未找到该书信息 bookId：%+v\n", bookId)
		return nil
	}
	return books[0]
}

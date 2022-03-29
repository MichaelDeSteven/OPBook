package model

import (
	"bytes"
	"strings"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"github.com/PuerkitoBio/goquery"
)

type Document struct {
	Id           int       `gorm:"primaryKey;auto;" json:"id"`
	Name         string    `gorm:"" json:"name"`     // 文档名称
	Identify     string    `gorm:"" json:"identify"` // 文档唯一标识
	BookId       int       `gorm:"" json:"book_id"`  // 书籍Id
	ParentId     int       `gorm:"" json:"parent_id"`
	OrderSort    int       `gorm:"" json:"order_sort"`
	Release      string    `gorm:"" json:"release"`                         // 发布后的Html格式内容
	CommentCount int       `gorm:"" json:"comment_count"`                   // 评论人数
	ViewCount    int       `gorm:"" json:"view_count"`                      // 文档被阅读次数
	UserId       int       `gorm:"" json:"user_id"`                         // 创建的用户id
	Version      int64     `gorm:"default:0;" json:"version"`               // 版本
	CreateTime   time.Time `gorm:"autoCreateTime" json:"create_time"`       // 创建时间
	ModifyTime   time.Time `gorm:"autoUpdateTime:milli" json:"modify_time"` // 更新时间
	IsDeleted    int       `gorm:"default:0;" json:"is_deleted"`            // 逻辑删除（0:可用 1:不可用）
	Markdown     string    `gorm:"-" json:"markdown"`
}

func (d *Document) TableName() string {
	return "opbook_doc"
}

func NewDocument() *Document {
	return &Document{}
}

// 插入和更新文档.
// 存在文档id或者文档标识，则表示更新文档内容
func (d *Document) InsertOrUpdate() (id int, err error) {
	id = d.Id
	d.Name = strings.TrimSpace(d.Name)
	if id > 0 {
		tx := global.DB.Updates(d)
		err = tx.Error
		return
	}

	var dd *Document = NewDocument()
	global.DB.Select("id").Where("is_deleted = ?", 0).Where("identify = ?", d.Identify).Where("book_id = ?", d.BookId).First(dd)
	if dd.Id == 0 {
		tx := global.DB.Create(d)
		err = tx.Error
		id = d.Id
		NewBook().ResetDocumentNumber(d.BookId)
	} else {
		id = dd.Id
		tx := global.DB.Where("id = ?", dd.Id).Updates(d)
		err = tx.Error
	}
	return
}

// 发布文档内容为HTML
func (d *Document) ReleaseContent(bookId int) {
	// 加锁
	utils.BooksRelease.Set(bookId)
	defer utils.BooksRelease.Delete(bookId)

	var (
		// releaseTime = time.Now() // 发布的时间戳
		book       = NewBook()
		docs       []Document
		ModelStore = NewDocumentStore()
	)
	global.DB.Where("is_deleted = ?", 0).Where("id = ?", bookId).First(&book)

	// 全部重新发布。查询该书籍的所有文档id
	global.DB.Select("id", "identify", "modify_time").Where("book_id = ?", bookId).Where("is_deleted = ?", 0).Find(&docs)

	docMap := make(map[string]bool)
	for _, item := range docs {
		docMap[item.Identify] = true
	}
	for _, item := range docs {
		ds, err := ModelStore.GetById(item.Id)
		if err != nil {
			global.LOG.Sugar().Error(err)
			continue
		}

		if item.Release = strings.TrimSpace(ds.Content); !utils.Empty(item.Release) {
			if docQuery, err := goquery.NewDocumentFromReader(bytes.NewBufferString(item.Release)); err == nil {
				if html, err := docQuery.Html(); err == nil {
					item.Release = strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(html), "<html><head></head><body>"), "</body></html>")
				}
			}
		}
		if _, err = item.InsertOrUpdate(); err != nil {
			global.LOG.Sugar().Error(err)
		}
	}

	client := NewElasticSearchClient()
	client.RebuildAllIndex(bookId)
}

// 根据书籍ID查询文档一级目录.
func (m *Document) GetMenuTop(bookId int) (docs []*Document) {
	var docsAll []*Document
	cols := []string{"id", "name", "parent_id", "book_id", "identify"}
	global.DB.Select(cols).Where("book_id = ?", bookId).Where("is_deleted = ?", 0).Where("parent_id = ?", 0).Order("order_sort, id").Limit(5000).Find(&docsAll)

	// 以"."开头的文档标识，不在阅读目录显示
	for _, doc := range docsAll {
		if !strings.HasPrefix(doc.Identify, ".") {
			docs = append(docs, doc)
		}
	}
	return
}

func (d *Document) Get(bookId int, docIdentify string) (doc *Document) {
	global.DB.Where("book_id = ?", bookId).Where("is_deleted = ?", 0).Where("identify = ?", docIdentify).First(&doc)
	return
}

func (d *Document) getDocsByBookId(bookId int, fields ...string) (docs []*Document) {
	global.DB.Model(d).Select(fields).Where("is_deleted = ?", 0).Where("book_id", bookId).Order("order_sort, identify").Limit(5000).Find(&docs)
	return
}

//分页查询文档
func (m *Document) FindToPager(pageIndex, pageSize, bookId int, fields ...string) (docs []*Document, err error) {

	offset := (pageIndex - 1) * pageSize

	tx := global.DB.Select(fields).Limit(pageSize).Offset(offset).Find(&docs)
	err = tx.Error
	return
}

func (d *Document) GetDocumentView(docId int) int {
	var model Book
	global.DB.Select("view_count").Where("is_deleted = ?", 0).Where("id = ?", docId).Find(&model)
	return model.ViewCount
}

func (d *Document) SetDocumentView(docId int, viewCount int) {
	global.DB.Update("view_count", viewCount).Where("is_deleted = ?", 0).Where("id = ?", docId)
}

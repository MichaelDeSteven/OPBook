package model

import (
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

// 文档结果
type DocResult struct {
	DocId        int       `json:"doc_id"`
	DocName      string    `json:"doc_name"`
	Identify     string    `json:"identify"`   // Identify 文档唯一标识
	Release      string    `json:"release"`    // Release 发布后的Html格式内容.
	ViewCount    int       `json:"view_count"` // 书籍被浏览次数
	CreateTime   time.Time `json:"create_time"`
	BookId       int       `json:"book_id"`
	BookIdentify string    `json:"book_identify"`
	BookName     string    `json:"book_name"`
}

func NewDocResult() *DocResult {
	return &DocResult{}
}

func (d *DocResult) GetDocsById(id []int) (docs []*DocResult) {
	if len(id) == 0 {
		return
	}
	global.DB.Table("opbook_book b").Select("d.id as doc_id, d.name as doc_name, d.identify, d.view_count, d.create_time, b.id as book_id, b.identify as book_identify, d.release, b.name as book_name").Joins("left join opbook_doc d on d.book_id = b.id").Where("b.is_deleted = ?", 0).Where("d.id in ?", id).Find(&docs)
	return
}

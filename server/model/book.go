package model

import (
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

// Book
type Book struct {
	Id                int       `gorm:"primaryKey;auto;" json:"id"`
	Name              string    `gorm:"" json:"name"`                            // 书籍名称
	Identify          string    `gorm:"" json:"identify"`                        // 书籍唯一标识
	OrderIndex        int       `gorm:"default:0;" json:"order_index"`           // 推荐权重越高, 数字越大排序越靠前
	Pin               int       `gorm:"default:0;" json:"pin"`                   // pin值，用于首页固定显示
	Description       string    `gorm:"" json:"description"`                     // 书籍描述
	Label             string    `gorm:"" json:"label"`                           // 书籍标签
	Private           int       `gorm:"default:0;" json:"private"`               // 书籍私有: 0:公开 1:私有
	PrivateToken      string    `gorm:"" json:"private_token"`                   // 当书籍是私有时的访问Token
	DocCount          int       `gorm:"" json:"doc_count"`                       // 包含文档数量
	Cover             string    `gorm:"" json:"cover"`                           // 封面地址
	CommentCount      int       `gorm:"" json:"comment_count"`                   // 评论人数
	ScoreCount        int       `gorm:"" json:"score_count"`                     // 评分人数
	Score             int       `gorm:"default:40;" json:"score"`                // 书籍评分，默认40，即4.0星
	ViewCount         int       `gorm:"" json:"view_count"`                      // 书籍被阅读次数
	CollectCount      int       `gorm:"" json:"collect_count"`                   // 书籍被收藏次数
	Author            string    `gorm:"" json:"author"`                          // 原作者，即来源
	AuthorURL         string    `gorm:"" json:"author_url"`                      // 原作者链接，即来源链接
	Lang              string    `gorm:"default:zh;" json:"lang"`                 // 语言
	UserId            int       `gorm:"" json:"user_id"`                         // 持有书籍的用户id
	Version           int64     `gorm:"default:0;" json:"version"`               // 版本
	CreateTime        time.Time `gorm:"autoCreateTime" json:"create_time"`       // 创建时间
	ModifyTime        time.Time `gorm:"autoUpdateTime:milli" json:"modify_time"` // 更新时间
	ReleaseTime       time.Time `gorm:"autoCreateTime" json:"release_time"`      // 书籍发布时间
	GenerateTime      time.Time `gorm:"autoCreateTime" json:"generate_time"`     // 下载文档生成时间
	LastClickGenerate time.Time `gorm:"autoCreateTime" json:"-"`                 // 上次点击上传文档的时间，用于显示频繁点击浪费服务器硬件资源的情况
	IsDeleted         int       `gorm:"default:0;" json:"is_deleted"`            // 逻辑删除（0:可用 1:不可用）
}

func (b *Book) TableName() string {
	return "opbook_book"
}

func NewBook() *Book {
	return &Book{}
}

func (b *Book) FindByIdentify(identify string) (book *Book) {
	global.DB.Where("is_deleted = ?", 0).Where("identify = ?", identify).First(&book)
	return book
}

func (b *Book) CheckIdentifyIsExist(identify string) bool {
	var count int64
	var model Book
	global.DB.Model(model).Where("is_deleted = ?", 0).Where("identify = ?", identify).Count(&count)
	return count > 0
}

func (b *Book) Add(book *Book) error {
	tx := global.DB.Create(book)
	return tx.Error
}

func (m *Book) HasProjectAccess(identify string, memberId int) bool {
	return true
}

//分页查询指定用户的书籍
//按照最新的进行排序
func (m *Book) FindToPager(pageIndex, pageSize, userId int) (books []Book, totalCount int64, err error) {
	var model Book
	global.DB.Model(model).Where("is_deleted = ?", 0).Count(&totalCount)

	offset := (pageIndex - 1) * pageSize

	tx := global.DB.Order("id desc").Limit(pageSize).Offset(offset).Find(&books)
	err = tx.Error
	return
}

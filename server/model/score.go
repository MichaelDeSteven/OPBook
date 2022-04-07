package model

import (
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

type Score struct {
	Id         int       `gorm:"primaryKey;auto;" json:"id"`
	UserId     int       `gorm:"" json:"user_id"`                   // 用户id
	BookId     int       `gorm:"" json:"book_id"`                   // 书籍id
	Score      int       `gorm:"default:40;" json:"score"`          // 书籍评分，默认40，即4.0星
	CreateTime time.Time `gorm:"autoCreateTime" json:"create_time"` // 创建时间
}

func (s *Score) TableName() string {
	return "opbook_score"
}

func NewScore() *Score {
	return &Score{}
}

// 添加评分
func (this *Score) AddScore() (err error) {
	tx := global.DB.Create(this)
	err = tx.Error
	return
}

// 获取评分
func (this *Score) GetScore(bid, uid int) (s *Score) {
	global.DB.Where("book_id = ?", bid).Where("user_id = ?", uid).First(&this)
	return this
}

package model

import "github.com/MichaelDeSteven/OPBook/server/global"

type Star struct {
	Id        int `gorm:"primaryKey;auto;" json:"id"`
	UserId    int `gorm:"" json:"user_id"`              // 用户id
	BookId    int `gorm:"" json:"book_id"`              // 书籍id
	IsDeleted int `gorm:"default:0;" json:"is_deleted"` // 逻辑删除（0:已收藏 1:未收藏）
}

func (s *Star) TableName() string {
	return "opbook_star"
}

func NewStar() *Star {
	return &Star{}
}

// 获取
func (s *Star) GetStarbyUserIdAndBookId(userId, bookId int) *Star {
	global.DB.Where("user_id = ?", userId).Where("book_id = ?", bookId).First(s)
	return s
}

// 收藏
func (s *Star) Star(userId, bookId int, exist bool) {
	if exist {
		global.DB.Model(s).Where("book_id = ?", bookId).Where("user_id = ?", userId).Update("is_deleted", s.IsDeleted)
	} else {
		s.UserId, s.BookId = userId, bookId
		global.DB.Create(s)
	}
}

// 取消收藏
func (s *Star) UnStar(userId, bookId int) {
	global.DB.Model(s).Where("book_id = ?", bookId).Where("user_id = ?", userId).Update("is_deleted", 1)
}

// 查询用户是否已收藏
func (s *Star) IsStar(userId, bookId int) bool {
	global.DB.Where("is_deleted = ?", 0).Where("user_id = ?", userId).Where("book_id = ?", bookId).First(s)
	return s.Id > 0
}

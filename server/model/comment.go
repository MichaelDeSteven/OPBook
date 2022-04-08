package model

import (
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

type Comment struct {
	Id          int       `gorm:"primaryKey;auto;" json:"id"`
	UserId      int       `gorm:"" json:"user_id"`                    // 用户id
	BookId      int       `gorm:"" json:"book_id"`                    // 书籍id
	Content     string    `gorm:"" json:"content"`                    // 评论内容
	LikeCount   int       `gorm:"" json:"like_count"`                 // 点赞数
	CommentId   int       `gorm:"" json:"comment_id"`                 // 被回复评论id
	CommentTime time.Time `gorm:"autoCreateTime" json:"comment_time"` // 评论时间
}

type CommentResult struct {
	Id            int       `json:"id"`
	UserId        int       `json:"user_id"`        // 用户id
	Nickname      string    `json:"nickname"`       // 昵称
	Avatar        string    `json:"avatar"`         // 头像url
	Content       string    `json:"content"`        // 评论内容
	LikeCount     int       `json:"like_count"`     // 点赞数
	CommentId     int       `json:"comment_id"`     // 被回复评论id
	ReplyNickname string    `json:"reply_nickname"` // 被回复昵称
	ReplyContent  string    `json:"reply_content"`  // 被回复评论
	CommentTime   time.Time `json:"comment_time"`   // 评论时间
	IsLike        bool      `json:"is_like"`        // 是否点赞
}

func (c *Comment) TableName() string {
	return "opbook_comment"
}

func NewComment() *Comment {
	return &Comment{}
}

func NewCommentResult() *CommentResult {
	return &CommentResult{}
}

// 添加评论
func (this *Comment) AddComment() (err error) {
	tx := global.DB.Create(this)
	err = tx.Error
	return
}

func (this *Comment) GetCommentByBookId(bookId int) (res []*CommentResult) {
	global.DB.Table("opbook_user u").Select("u.id as user_id, u.nickname as nickname, u.avatar as avatar, c.content as content, c.comment_id as comment_id, c.comment_time as comment_time, c.id as id").Joins("left join opbook_comment c on u.id = c.user_id").Where("c.is_deleted = ?", 0).Where("c.book_id = ?", bookId).Find(&res)
	return
}

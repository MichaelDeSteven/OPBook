package model

import (
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

type Message struct {
	Id             int       `gorm:"primaryKey;auto;" json:"id"`
	Content        string    `gorm:"" json:"content"`                   // 评论内容
	FromId         int       `gorm:"" json:"from_id"`                   // 发送方Id
	ToId           int       `gorm:"" json:"to_id"`                     // 接收方Id
	ConversationId string    `gorm:"" json:"conversation_id"`           // 会话Id
	CreateTime     time.Time `gorm:"autoCreateTime" json:"create_time"` // 创建时间
}

func (m *Message) TableName() string {
	return "opbook_message"
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) Add() (err error) {
	tx := global.DB.Create(m)
	err = tx.Error
	return
}

func (m *Message) GetConversationById(conversationId string) (list []*Message) {
	global.DB.Where("conversation_id = ?", conversationId).Find(&list)
	return
}

func (m *Message) GetUserList(uid int) (list []*Message) {
	global.DB.Select("from_id, to_id").Where("from_id = ?", uid).Or("to_id = ?", uid).Group("conversation_id").Find(&list)
	return
}

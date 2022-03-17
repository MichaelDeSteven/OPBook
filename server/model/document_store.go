package model

import (
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

type DocumentStore struct {
	DocumentId int       `gorm:"primaryKey;"` // 对应Document中的id
	Markdown   string    `gorm:""`            // markdown内容
	Content    string    `gorm:""`            // 文本内容
	ModifyTime time.Time `gorm:"-"`           // 更新时间
}

func (d *DocumentStore) TableName() string {
	return "opbook_doc_store"
}

func NewDocumentStore() *DocumentStore {
	return &DocumentStore{}
}

// 插入或者更新
func (this *DocumentStore) InsertOrUpdate(ds *DocumentStore, fields ...string) (err error) {
	ds.ModifyTime = time.Now().Add(1 * time.Second)

	var one *DocumentStore = NewDocumentStore()
	global.DB.Select("document_id").Where("document_id = ?", ds.DocumentId).First(one)

	if one.DocumentId > 0 {
		if len(fields) > 0 {
			var updateFields []string
			withoutUpdatedAt := false
			for _, field := range fields {
				if field == "-updated_at" || field == "-UpdatedAt" {
					withoutUpdatedAt = true
					continue
				}

				if field == "updated_at" || field == "UpdatedAt" {
					continue
				}
				updateFields = append(updateFields, field)
			}

			fields = updateFields

			if withoutUpdatedAt == false {
				fields = append(fields, "modify_time")
			}
		}
		tx := global.DB.Model(ds).Where("document_id", one.DocumentId).Select(fields).Updates(ds)
		err = tx.Error
	} else {
		tx := global.DB.Create(ds)
		err = tx.Error
	}
	return
}

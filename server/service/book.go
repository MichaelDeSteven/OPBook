package service

import (
	"errors"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/model"
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

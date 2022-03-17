package utils

import (
	"errors"
	"os"
	"strings"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/PuerkitoBio/goquery"
	"go.uber.org/zap"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			global.LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.LOG.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}

// 从md的html文件中提取文章标题（从h1-h6）
func ParseTitleFromMdHtml(html string) (title string) {
	hTag := []string{"h1", "h2", "h3", "h4", "h5", "h6"}
	if doc, err := goquery.NewDocumentFromReader(strings.NewReader(html)); err == nil {
		for _, tag := range hTag {
			if title = strings.TrimSpace(doc.Find(tag).First().Text()); title != "" {
				return title
			}
		}
	} else {
		global.LOG.Sugar().Errorf("%+v\n", err)
	}
	return "空标题文档"
}

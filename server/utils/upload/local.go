package upload

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/utils"
	"go.uber.org/zap"
)

type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := path.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.CONFIG.Local.Path, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.CONFIG.Local.Path + "/" + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return p, filename, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	if err := os.Remove(key); err != nil {
		return errors.New("本地文件删除失败, err:" + err.Error())
	}
	return nil
}

func (*Local) UploadFileByPath(src string, fileName, ext string) (string, string, error) {
	f, err := os.OpenFile(src, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return "", "", err
	}
	defer f.Close()
	target := filepath.Join("./", global.CONFIG.Local.Path, global.CONFIG.Local.Avator, time.Now().Format("200601"))
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(target, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	dst := filepath.Join(target, fileName+ext)
	out, createErr := os.Create(dst)
	if createErr != nil {
		global.LOG.Error("function os.Create() Filed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close()             // 创建文件 defer 关闭
	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.LOG.Error("function io.Copy() Filed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return dst, fileName + ext, nil
}
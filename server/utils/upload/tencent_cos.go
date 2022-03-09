package upload

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
)

type TencentCOS struct{}

// UploadFile upload file to COS
func (*TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		global.LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), global.CONFIG.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return global.CONFIG.TencentCOS.BaseURL + "/" + global.CONFIG.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}

func (*TencentCOS) UploadFileByPath(src string, fileName, ext string) (string, string, error) {
	client := NewClient()
	f, err := os.OpenFile(src, os.O_SYNC|os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", "", err
	}
	defer f.Close()
	target := filepath.Join("./", global.CONFIG.Local.Path, time.Now().Format("200601"))
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(target, os.ModePerm)
	if mkdirErr != nil {
		global.LOG.Error("function os.MkdirAll() Filed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	_, err = client.Object.Put(context.Background(), global.CONFIG.TencentCOS.PathPrefix+"/"+fileName+ext, f, nil)
	if err != nil {
		panic(err)
	}
	return global.CONFIG.TencentCOS.BaseURL + "/" + global.CONFIG.TencentCOS.PathPrefix + "/" + fileName + ext, fileName, nil
}

// DeleteFile delete file form COS
func (*TencentCOS) DeleteFile(url string) error {
	client := NewClient()
	strs := strings.Split(url, "/")
	name := global.CONFIG.TencentCOS.PathPrefix + "/" + strs[len(strs)-1]
	_, err := client.Object.Delete(context.Background(), name)
	if err != nil {
		global.LOG.Error("function bucketManager.Delete() Filed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() Filed, err:" + err.Error())
	}
	return nil
}

// NewClient init COS client
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + global.CONFIG.TencentCOS.Bucket + ".cos." + global.CONFIG.TencentCOS.Region + ".myqcloud.com")
	baseURL := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseURL, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.CONFIG.TencentCOS.SecretID,
			SecretKey: global.CONFIG.TencentCOS.SecretKey,
		},
	})
	return client
}

package upload

import (
	"mime/multipart"

	"github.com/MichaelDeSteven/OPBook/server/global"
)

// OSS 对象存储接口
type OSS interface {
	UploadFile(file *multipart.FileHeader) (string, string, error)
	UploadFileByPath(src string, fileName, ext string) (string, string, error)
	DeleteFile(key string) error
}

// NewOss OSS的实例化方法
// Author [SliverHorn](https://github.com/SliverHorn)
// Author [ccfish86](https://github.com/ccfish86)
func NewOss() OSS {
	switch global.CONFIG.System.OssType {
	case "local":
		return &Local{}
	case "tencent-cos":
		return &TencentCOS{}
	default:
		return &Local{}
	}
}

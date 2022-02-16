package system

import (
	"github.com/MichaelDeSteven/OPBook/server/config"
	"github.com/MichaelDeSteven/OPBook/server/global"
)

//@author: [MichaelDeSteven](https://github.com/MichaelDeSteven)
//@function: GetSystemConfig
//@description: 读取配置文件
//@return: err error, conf config.Server

type SystemConfigService struct{}

func (systemConfigService *SystemConfigService) GetSystemConfig() (err error, conf config.Server) {
	return nil, global.CONFIG
}

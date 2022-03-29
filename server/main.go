package main

import (
	"github.com/MichaelDeSteven/OPBook/server/core"
	"github.com/MichaelDeSteven/OPBook/server/global"
	"github.com/MichaelDeSteven/OPBook/server/initialize"
)

func main() {
	global.VP = core.Viper()      // 初始化Viper
	global.DB = initialize.Gorm() // gorm连接数据库
	global.LOG = core.Zap()       // 初始化zap日志库
	initialize.Redis()
	initialize.Routers().Start()
}

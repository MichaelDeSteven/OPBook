package initialize

import (
	"github.com/MichaelDeSteven/OPBook/server/global"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author MichaelDeSteven
func initGorm() {
	switch global.CONFIG.System.DbType {
	case "mysql":
		global.DB = GormMysql()
	default:
		global.DB = GormMysql()
	}
}

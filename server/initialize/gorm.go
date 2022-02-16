package initialize

import (
	"github.com/MichaelDeSteven/OPBook/server/global"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author MichaelDeSteven
func Gorm() *gorm.DB {
	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

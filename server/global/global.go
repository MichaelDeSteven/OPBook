package global

import (
	"github.com/MichaelDeSteven/OPBook/server/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB                  *gorm.DB
	CONFIG              config.Server
	VP                  *viper.Viper
	LOG                 *zap.Logger
	Concurrency_Control = &singleflight.Group{}
)

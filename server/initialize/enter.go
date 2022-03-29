package initialize

func Init() {
	initViper() // 初始化Viper
	initGorm()  // gorm连接数据库
	initLog()   // 初始化zap日志库
	initTimer()
	initRedis()
	initRouters().Start()
}

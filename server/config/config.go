package config

type Server struct {
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// system
	System System `mapstructure:"system" json:"system" yaml:"system"`
	// Zap log
	Zap Zap `mapstructure:"zap" json:"zap" yaml:"zap"`
	// 鉴权
	JWT JWT `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}

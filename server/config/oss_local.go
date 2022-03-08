package config

type Local struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                         // 资源所在项目根下的目录名
	Avator       string `mapstructure:"avator" json:"avator" yaml:"avator"`                   // 头像相对目录目录
	Book         string `mapstructure:"book" json:"book" yaml:"book"`                         // 书籍相对目录目录
	RelativePath string `mapstructure:"relativePath" json:"relativePath" yaml:"relativePath"` // 相对路径
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`                         // ip+端口
}

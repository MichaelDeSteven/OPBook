package config

type Local struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                         // 资源所在项目根下的目录名
	RelativePath string `mapstructure:"relativePath" json:"relativePath" yaml:"relativePath"` // 相对路径
}

package config

type Timer struct {
	Start bool   `mapstructure:"start" json:"start" yaml:"start"` // 是否启用
	Spec  string `mapstructure:"spec" json:"spec" yaml:"spec"`    // CRON表达式
}

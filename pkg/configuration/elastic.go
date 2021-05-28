package configuration

/*
Elastic elasticsearch 数据库连接配置

示例：
user: elasticsearch
password: xxx
url: http://127.0.0.1:9200
max_open_conns: 1000
max_idle_conns: 500
*/
type Elastic struct {
	User         string `mapstructure:"user" validate:"required"`
	Password     string `mapstructure:"password" validate:"required"`
	URL          string `mapstructure:"url" validate:"required,url"`
	MaxOpenConns int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
	Prefix       string `mapstructure:"prefix" validate:"-"`

	// Debug 当值位 true 时在标准输出打印 elasticsearch 真实请求语句及响应结果
	Debug bool `mapstructure:"debug" validate:"-"`
}

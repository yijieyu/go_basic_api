package configuration

import "time"

/*
HTTP http 服务相关配置

示例：
http:
  host: 127.0.0.1
  port: 8080
  mode: debug
*/
type HTTP struct {
	// Host http 服务启动监听的主机名
	Host string `mapstructure:"host" validate:"required,ip"`

	// Port http 服务启动监听的端口
	Port int `mapstructure:"port" validate:"required"`

	// Mode http 服务运行级别
	Mode string `mapstructure:"mode" validate:"required,eq=release|eq=test|eq=debug"`

	// StopTimeout 优雅关闭超时时间
	StopTimeout time.Duration `mapstructure:"stop_timeout"`
}

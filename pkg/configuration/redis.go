package configuration

/*
Redis redis 数据库配置

示例：
host: 127.0.0.1
port: 6379
password: auth
db: 0
idle_timeout: 1
max_open_conns: 1000
max_idle_conns: 200
*/
type Redis struct {
	Host         string `mapstructure:"host" validate:"required"`
	Port         int    `mapstructure:"port" validate:"required"`
	Password     string `mapstructure:"password"`
	DB           int    `mapstructure:"db"`
	IdleTimeout  int    `mapstructure:"idle_timeout" validator:"required"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns" validator:"required,min=1,ltefield=MaxOpenConns"`
}

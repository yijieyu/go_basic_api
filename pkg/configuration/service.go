package configuration

/*
Service 服务配置

示例：
appname: app_api
host: 127.0.0.1
public_ip: 127.0.0.1
environment: testing
*/
type Service struct {
	// AppName 应用名称
	AppName string `mapstructure:"appname" validate:"required"`

	// Host 内网ip
	Host string `mapstructure:"host" validate:"required,ip"`

	// PublicIP 公网ip
	PublicIP string `mapstructure:"public_ip" validate:"ip"`

	// Environment 运行环境
	Environment string `mapstructure:"environment" validate:"required,eq=testing|eq=debug|eq=prod"`

	Address string `mapstructure:"address" validate:"required"`
}

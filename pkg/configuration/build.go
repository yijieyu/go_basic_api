package configuration

/*
Build 应用程序编译信息

示例：
go_version: go_version
version: version_number
build_date: build_date
*/
type Build struct {
	GoVersion string `mapstructure:"go_version" validate:"required"`
	Version   string `mapstructure:"version" validate:"required"`
	BuildDate string `mapstructure:"build_date" validate:"required"`
}

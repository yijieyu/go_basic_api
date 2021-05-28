package configuration

/*
DB 关系型数据库配置

示例：
data_source_name: username:password@tcp(127.0.0.1:3306)/dbname?parseTime=true&charset=utf8&loc=Local
max_open_conns: 2
max_idle_conns: 2
*/
type DB struct {
	// please refer to database drivers in github
	// e.g. mysql driver: https://github.com/go-sql-driver/mysql#dsn-data-source-name
	DataSourceName string `mapstructure:"data_source_name" validate:"required,dsn"`
	MaxOpenConns   int    `mapstructure:"max_open_conns" validate:"required,min=1"`
	MaxIdleConns   int    `mapstructure:"max_idle_conns" validate:"required,min=1,ltefield=MaxOpenConns"`
}

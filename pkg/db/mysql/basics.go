package mysql

import "context"

type Basics interface {
	Ping() error
	PingContext(ctx context.Context) error
	DriverName() string
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	MapperFunc(mf func(string) string)
}

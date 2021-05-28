package mysql

import (
	"github.com/yijieyu/go_basic_api/pkg/db/mysql"
)

type Storage struct {
	c *mysql.Client
}

func NewStorage(client *mysql.Client) *Storage {

	return &Storage{
		c: client,
	}
}

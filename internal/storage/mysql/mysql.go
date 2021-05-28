package mysql

import (
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/db/mysql"
)

type Storage struct {
	c *mysql.Client
}

func NewStorage(client *mysql.Client) *Storage {

	return &Storage{
		c: client,
	}
}

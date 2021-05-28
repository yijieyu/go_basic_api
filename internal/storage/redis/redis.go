package redis

import (
	"github.com/yijieyu/go_basic_api/pkg/db/redis"
)

type Storage struct {
	c *redis.Client
}

func NewStorage(client *redis.Client) *Storage {
	return &Storage{
		c: client,
	}
}

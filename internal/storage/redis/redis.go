package redis

import (
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/db/redis"
)

type Storage struct {
	c *redis.Client
}

func NewStorage(client *redis.Client) *Storage {
	return &Storage{
		c: client,
	}
}

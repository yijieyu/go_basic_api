package elasticsearch

import (
	"github.com/thinkeridea/go-extend/pool"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/db/elasticsearch"
)

type Storage struct {
	c      *elasticsearch.Client
	prefix string

	pool pool.BufferPool
}

func NewStorage(client *elasticsearch.Client, prefix string) *Storage {
	return &Storage{
		c:      client,
		prefix: prefix,
		pool:   pool.GetBuff4096(),
	}
}

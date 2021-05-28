package etcd

import (
	"errors"

	"github.com/coreos/etcd/clientv3"
)

var (
	NotFoundKeyErr = errors.New("etcd not found key")
)

type Configuration interface {
}

type configuration struct {
	prefix string
	*clientv3.Client
}

func NewConfiguration(db *clientv3.Client, prefix string) Configuration {
	return &configuration{Client: db, prefix: prefix}
}

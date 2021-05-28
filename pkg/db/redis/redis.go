package redis

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/yijieyu/go_basic_api/pkg/configuration"
)

type Client struct {
	master []*redis.Pool
	slave  []*redis.Pool
}

func New(master []configuration.Redis, slave ...configuration.Redis) *Client {
	db := &Client{
		master: make([]*redis.Pool, len(master)),
		slave:  make([]*redis.Pool, len(slave)),
	}
	for i := range master {
		pool := &redis.Pool{
			MaxIdle:     master[i].MaxIdleConns,
			MaxActive:   master[i].MaxOpenConns,
			IdleTimeout: time.Duration(master[i].IdleTimeout) * time.Second,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
			Dial: db.connect(master[i]),
		}
		db.master[i] = pool
	}

	for i := range slave {
		pool := &redis.Pool{
			MaxIdle:     slave[i].MaxIdleConns,
			MaxActive:   slave[i].MaxOpenConns,
			IdleTimeout: time.Duration(slave[i].IdleTimeout) * time.Second,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
			Dial: db.connect(slave[i]),
		}
		db.slave[i] = pool
	}

	if len(db.slave) == 0 {
		db.slave = db.master
	}

	return db
}

func (db *Client) connect(config configuration.Redis) func() (redis.Conn, error) {
	return func() (redis.Conn, error) {
		return redis.Dial(
			"tcp",
			fmt.Sprintf("%s:%d", config.Host, config.Port),
			redis.DialDatabase(config.DB),
			redis.DialPassword(config.Password),
		)
	}
}

func (db *Client) Writer() redis.Conn {
	return db.master[rand.Intn(len(db.master))].Get()
}

func (db *Client) Reader() redis.Conn {
	return db.slave[rand.Intn(len(db.slave))].Get()
}

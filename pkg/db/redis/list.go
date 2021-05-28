package redis

import (
	"github.com/garyburd/redigo/redis"
)

type List interface {
	Push(key, value string) error
	POP(key string) (string, error)
	LPush(key, value string) error
}

// Push 压入一个元素到队列末尾
func (db *Client) Push(key, value string) error {
	c := db.Writer()
	defer c.Close()

	_, err := c.Do("RPUSH", key, value)
	return err
}

// POP 获取队列的一个元素
func (db *Client) POP(key string) (string, error) {
	c := db.Writer()
	defer c.Close()

	ss, err := redis.ByteSlices(c.Do("BLPOP", key, 0))
	if err != nil || len(ss) < 2 {
		return "", err
	}

	return string(ss[1]), err
}

// LPush 业务处理失败，往队列头压入元素
func (db *Client) LPush(key, value string) error {
	c := db.Writer()
	defer c.Close()

	_, err := c.Do("LPUSH", key, value)
	return err
}

// RPOP 获取队列的一个元素
func (db *Client) RPOP(key string) (string, error) {
	c := db.Writer()
	defer c.Close()

	ss, err := redis.ByteSlices(c.Do("BRPOP", key, 0))
	if err != nil || len(ss) < 2 {
		return "", err
	}

	return string(ss[1]), err
}

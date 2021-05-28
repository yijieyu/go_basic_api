package redis

import (
	"github.com/garyburd/redigo/redis"
)

func (db *Client) Get(key string) (string, error) {
	c := db.Reader()
	defer c.Close()

	v, err := redis.String(c.Do("get", key))
	if err == redis.ErrNil {
		v = ""
		err = nil
	}

	return v, err
}

func (db *Client) Set(key, value string, expire int) error {
	c := db.Writer()
	defer c.Close()

	_, err := c.Do("SETEX", key, expire, value)
	return err
}

func (db *Client) Del(key string) error {
	c := db.Writer()
	defer c.Close()

	_, err := c.Do("DEL", key)
	return err
}

package app

import (
	"github.com/yijieyu/go_basic_api/internal/storage/elasticsearch"
	"github.com/yijieyu/go_basic_api/internal/storage/mysql"
	"github.com/yijieyu/go_basic_api/internal/storage/redis"
)

type Storage struct {
	mysql   *mysql.Storage
	redis   *redis.Storage
	elastic *elasticsearch.Storage
}

func (s *Storage) init(app *App) {
	s.mysql = mysql.NewStorage(app.DB.mysql)
	s.redis = redis.NewStorage(app.DB.redis)

}

func (s *Storage) MySql() *mysql.Storage {
	return s.mysql
}

func (s *Storage) Elastic() *elasticsearch.Storage {
	return s.elastic
}

func (s *Storage) Redis() *redis.Storage {
	return s.redis
}

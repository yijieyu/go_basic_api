package app

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/thinkeridea/go-extend/helper"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/db/elasticsearch"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/db/mysql"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/db/redis"
)

type DB struct {
	mysql   *mysql.Client
	etcd    *clientv3.Client
	elastic *elasticsearch.Client
	redis   *redis.Client
}

func (db *DB) init(app *App) {
	db.mysql = mysql.New(app.config.DB.Master, app.config.DB.Slave)
	db.etcd = helper.Must(clientv3.New(clientv3.Config{Endpoints: app.config.Etcd})).(*clientv3.Client)
	db.redis = redis.New(app.config.Redis.Master, app.config.Redis.Slave...)
}

func (db *DB) MySql() *mysql.Client {
	return db.mysql
}

func (db *DB) Etcd() *clientv3.Client {
	return db.etcd
}

func (db *DB) Elastic() *elasticsearch.Client {
	return db.elastic
}

func (db *DB) Redis() *redis.Client {
	return db.redis
}

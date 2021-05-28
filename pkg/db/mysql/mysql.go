package mysql

import (
	"bytes"
	"math/rand"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/thinkeridea/go-extend/helper"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mapper = reflectx.NewMapperFunc("db", sqlx.NameMapper)

type db interface {
	Basics
	reader
	writer
}

type Client struct {
	master []*gorm.DB
	slave  []*gorm.DB
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func New(master []configuration.DB, slave []configuration.DB) *Client {
	s := &Client{
		master: make([]*gorm.DB, len(master)),
		slave:  make([]*gorm.DB, len(slave)),
	}

	for i := range master {
		s.master[i] = helper.Must(gorm.Open(mysql.Open(master[i].DataSourceName), &gorm.Config{})).(*gorm.DB)
		//db.SetMaxIdleConns(master[i].MaxIdleConns)
		//db := sqlx.MustConnect("mysql", master[i].DataSourceName).Unsafe()
		//db.SetMaxOpenConns(master[i].MaxOpenConns)
	}

	for i := range slave {
		s.slave[i] = helper.Must(gorm.Open(mysql.Open(master[i].DataSourceName), &gorm.Config{})).(*gorm.DB)
		//db := sqlx.MustConnect("mysql", slave[i].DataSourceName).Unsafe()
		//db.SetMaxIdleConns(slave[i].MaxIdleConns)
		//db.SetMaxOpenConns(slave[i].MaxOpenConns)
	}

	if len(s.slave) == 0 {
		s.slave = s.master
	}

	return s
}

func (db *Client) Writer() *gorm.DB {
	return db.master[rand.Intn(len(db.master))]
}

func (db *Client) Reader() *gorm.DB {

	return db.slave[rand.Intn(len(db.slave))]
}

func GetTableFieldNames(t reflect.Type) string {
	buf := &bytes.Buffer{}
	for tagName := range mapper.TypeMap(t).Names {
		buf.WriteString(",`")
		buf.WriteString(tagName)
		buf.WriteString("`")
	}

	return buf.String()[1:]
}

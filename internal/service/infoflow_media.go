package service

import (
	"github.com/yijieyu/go_basic_api/internal/model/weimiao"
	"github.com/yijieyu/go_basic_api/internal/storage/mysql"
)

type InfoflowMediaService struct {
	db *mysql.Storage
}

func NewInfoflowMediaService(db *mysql.Storage) *InfoflowMediaService {

	return &InfoflowMediaService{
		db: db,
	}
}

func (i *InfoflowMediaService) Get(media weimiao.InfoflowMedia) *weimiao.InfoflowMedia {

	return i.db.GetInfoflowMedia(media)
}

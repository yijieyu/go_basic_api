package service

import (
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/model/weimiao"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/storage/mysql"
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

package mysql

import (
	"errors"

	"github.com/yijieyu/go_basic_api/internal/model/weimiao"
	"gorm.io/gorm"
)

func (db *Storage) GetInfoflowMedia(media weimiao.InfoflowMedia) (res *weimiao.InfoflowMedia) {
	c := db.c.Reader()
	result := c.Where(media).First(&res)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return
}

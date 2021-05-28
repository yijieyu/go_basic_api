package infoflow_media

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/model/weimiao"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/service"
)

type InfoflowMedia struct {
	service *service.InfoflowMediaService
}

func NewInfoflowMedia(service *service.InfoflowMediaService) *InfoflowMedia {
	return &InfoflowMedia{
		service: service,
	}
}

func (i *InfoflowMedia) Get(c *gin.Context) {

	m := weimiao.InfoflowMedia{}
	m.ID, _ = strconv.Atoi(c.Query("id"))

	c.JSON(http.StatusOK, gin.H{
		"data": i.service.Get(m),
	})
}

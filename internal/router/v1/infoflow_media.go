package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.weimiaocaishang.com/weimiao/base_api/app"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/server/infoflow_media"
)

func init() {
	router.Register("infoflow", InfoflowMedia)
}

func InfoflowMedia(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error {

	ad := infoflow_media.NewInfoflowMedia(app.Service.InfoflowMedia())
	g.GET("get", ad.Get)

	return nil
}

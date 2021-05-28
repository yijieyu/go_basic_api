package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.weimiaocaishang.com/weimiao/base_api/app"
	v1 "gitlab.weimiaocaishang.com/weimiao/base_api/internal/server/ad"
)

func init() {
	router.Register("ad", Ad)
}

func Ad(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error {

	ad := v1.NewAd(app.Service.Ad(), app.Conf().ThirdAdSwitch)
	g.GET("", ad.GetAd)

	return nil
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yijieyu/go_basic_api/app"
	v1 "github.com/yijieyu/go_basic_api/internal/server/ad"
)

func init() {
	router.Register("ad", Ad)
}

func Ad(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error {

	ad := v1.NewAd(app.Service.Ad(), app.Conf().ThirdAdSwitch)
	g.GET("", ad.GetAd)

	return nil
}

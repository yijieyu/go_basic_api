package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.weimiaocaishang.com/weimiao/base_api/app"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/router/group"
)

var router = group.New()

func Load(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error {
	return router.Load(app, g, mw...)
}

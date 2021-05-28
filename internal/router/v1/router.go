package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yijieyu/go_basic_api/app"
	"github.com/yijieyu/go_basic_api/internal/router/group"
)

var router = group.New()

func Load(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error {
	return router.Load(app, g, mw...)
}

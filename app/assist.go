package app

import (
	"github.com/gin-gonic/gin"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/debug"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/reload"
)

type Assist struct {
	debug  *debug.OnDebug
	reload *reload.OnReload
}

func (a *Assist) init(app *App) {
	a.debug = debug.New()
	a.reload = reload.New()

}

func (a *Assist) OnDebugRegister(name string, f debug.OnDebugFunc) {
	a.debug.Register(name, f)
}

func (a *Assist) OnReloadRegister(name string, f reload.OnReloadFunc, events ...string) {
	a.reload.Register(name, f, events...)
}

func (a *Assist) OnDebugFunc() func(cmd string, c *gin.Context) error {
	return a.debug.Run
}

func (a *Assist) OnReloadFunc() func(events ...string) error {
	return a.reload.Run
}

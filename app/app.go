package app

import (
	"sync"

	"github.com/thinkeridea/go-extend/helper"
	"github.com/yijieyu/go_basic_api/pkg/constant"
	"github.com/yijieyu/go_basic_api/pkg/errno"

	"github.com/yijieyu/go_basic_api/pkg/logger"
)

type App struct {
	config *Config
	env    string

	rw sync.RWMutex

	Assist
	DB
	Service
	Storage
	Oss
	Provider
}

func New(env string) *App {

	app := &App{
		env: env,
	}

	InitConfig(app)
	helper.Must(nil, logger.InitLog(app.config.Log))

	app.Assist.init(app)
	app.DB.init(app)
	app.Storage.init(app)
	app.Service.init(app)
	app.Oss.init(app)

	errno.DebugMode = app.config.Service.Environment == "testing"
	constant.JwtSecretKey = app.config.Jwt.SecretKey

	return app
}

func (app *App) Name() string {
	return app.config.Service.AppName
}

func (app *App) Conf() Config {
	return *app.config
}

func (app *App) Environment() string {
	return app.config.Service.Environment
}

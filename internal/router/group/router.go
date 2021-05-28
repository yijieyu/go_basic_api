package group

import (
	"github.com/gin-gonic/gin"
	"github.com/yijieyu/go_basic_api/app"
)

type RoutePlugin func(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error

type Group struct {
	plugins map[string]RoutePlugin
}

func New() *Group {
	return &Group{plugins: make(map[string]RoutePlugin, 10)}
}

func (r *Group) Register(name string, f RoutePlugin) {
	if _, ok := r.plugins[name]; ok {
		panic("duplicate register routing components by " + name)
	}
	r.plugins[name] = f
}

func (r *Group) Load(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error {
	var err error
	for name, f := range r.plugins {
		if err = f(app, g.Group(name)); err != nil {
			return err
		}
	}

	return err
}

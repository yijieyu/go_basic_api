package debug

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type DebugPlugin interface {
	OnDebug(c *gin.Context) error
}

type OnDebugFunc func(c *gin.Context) error

type OnDebug struct {
	plugins map[string]OnDebugFunc
}

func New() *OnDebug {
	return &OnDebug{
		plugins: map[string]OnDebugFunc{},
	}
}

func (d *OnDebug) Register(name string, f OnDebugFunc) {
	if _, ok := d.plugins[name]; ok {
		panic("duplicate register debug components by " + name)
	}
	d.plugins[name] = f
}

func (d *OnDebug) Run(cmd string, c *gin.Context) error {
	command, ok := d.plugins[cmd]
	if !ok {
		return fmt.Errorf("not found debug %s command", cmd)
	}

	return command(c)
}

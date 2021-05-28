// Package router 路由管理组件
package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/thinkeridea/go-extend/exnet/exhttp/expprof"
	"github.com/yijieyu/go_basic_api/app"
	v1 "github.com/yijieyu/go_basic_api/internal/router/v1"
	api "github.com/yijieyu/go_basic_api/internal/server"
	"github.com/yijieyu/go_basic_api/pkg/middleware"
)

type LoadFunc func(app *app.App, g *gin.RouterGroup, mw ...gin.HandlerFunc) error

func Load(app *app.App, engine *gin.Engine, mw ...gin.HandlerFunc) error {

	engine.Use(
		middleware.RecoveryWithWriter(os.Stdout),
		middleware.RequestID,
		middleware.ClientIP,
		middleware.Logging(),
		middleware.CORS(),
	)

	engine.NoRoute(api.NotFound)
	engine.NoMethod(api.NotFound)

	// 负载均衡健康检查接口
	engine.HEAD("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 验证服务名称
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, app.Name())
	})

	{
		// 内部服务，只能通过内网调用
		interior := engine.Group("/", middleware.DropRequestFromPublicNetwork)
		assist := api.NewAssist()
		// 运行时设置日志级别
		interior.GET("/log/:level", assist.SetLogLevel(app.Conf().Log.Level))
		// go pprof 运行时性能分析
		interior.GET(expprof.RoutePrefix+"*cmd", gin.WrapF(expprof.ServeHTTP))
		// 服务简况状态统计
		metrics := promhttp.Handler()
		interior.GET("/metrics", func(c *gin.Context) {
			metrics.ServeHTTP(c.Writer, c.Request)
		})

		// debug 接口
		interior.GET("/debug/:cmd", assist.Debug(app.OnDebugFunc()))
		interior.GET("/reload", assist.Reload(app.OnReloadFunc()))

		// config 配置
		interior.GET("/config", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": app.Conf(),
			})
			return
		})
	}

	var err error
	for path, load := range map[string]LoadFunc{
		"/v1": v1.Load,
	} {
		err = load(app, engine.Group(path))
		if err != nil {
			return err
		}
	}
	return err
}

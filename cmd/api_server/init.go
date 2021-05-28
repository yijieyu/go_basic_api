package api_server

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/thinkeridea/go-extend/helper"
	"github.com/yijieyu/go_basic_api/app"
	"github.com/yijieyu/go_basic_api/internal/router"
	"github.com/yijieyu/go_basic_api/pkg/services/http"
	"github.com/yijieyu/go_basic_api/pkg/signals"
)

var (
	env string
)

var ServerCmd = &cobra.Command{
	Use:   "api",
	Short: "api 常驻服务",
	Long:  `api 常驻服务 面向接口`,
	Run: func(cmd *cobra.Command, args []string) {

		// 初始化程序

		srv := app.New(env)

		// new http framework
		engine := gin.New()
		gin.SetMode(srv.Conf().HTTP.Mode)
		helper.Must(nil, router.Load(srv, engine))

		/*======================== start defer ===============================*/
		defer http.OnServiceStop(srv.Conf().HTTP.StopTimeout)
		go http.Catch(func() {
			http.OnServiceStop(srv.Conf().HTTP.StopTimeout)
		}, signals.Get()...)
		/*========================= end defer ==============================*/

		// start http service
		http.Run(srv.Conf().HTTP, engine)
	},
}

func init() {
	ServerCmd.PersistentFlags().StringVarP(&env, "env", "e", "", "env 环境变量")

}

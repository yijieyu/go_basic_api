package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thinkeridea/go-extend/exnet"
	"github.com/yijieyu/go_basic_api/pkg/constant"
)

// DropRequestFromPublicNetwork drops all request from public network, status code 403
func DropRequestFromPublicNetwork(c *gin.Context) {
	ip := c.GetString(constant.HttpClientIP)
	if !exnet.HasLocalIPddr(ip) {
		logrus.WithFields(logrus.Fields{
			"ip":   ip,
			"path": c.Request.URL.Path,
		}).Info("drops request from public network")
		c.AbortWithStatus(http.StatusForbidden)
		return
	}
}

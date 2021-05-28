package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkeridea/go-extend/exnet"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/constant"
)

func ClientIP(c *gin.Context) {
	ip := exnet.ClientPublicIP(c.Request)
	if ip == "" {
		ip = exnet.ClientIP(c.Request)
	}

	c.Set(constant.HttpClientIP, ip)
}

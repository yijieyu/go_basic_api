package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotFound)
}

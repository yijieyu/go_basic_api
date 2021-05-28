package middleware

import (
	"time"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

// CORS allow cross domain resources sharing
func CORS() gin.HandlerFunc {
	config := cors.Config{}
	config.AllowedHeaders = []string{"Content-Type", "Api-Key", "Signature"}
	config.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD"}
	config.AbortOnError = true
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.MaxAge = time.Hour * 12
	return cors.New(config)
}

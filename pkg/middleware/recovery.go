package middleware

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/constant"
)

const k = 1 << 10

// RecoveryWithWriter returns a middleware that
// recovers from any panics and writes a 500 response code
func RecoveryWithWriter(out io.Writer) gin.HandlerFunc {
	loggerStd := log.New(out, "", log.LstdFlags)

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, true)
				buf := make([]byte, 4*k)
				n := runtime.Stack(buf, false)

				loggerStd.Printf("[Recovery] panic recovered:\n%s\n%s\n%s\n", httpRequest, err, buf[:n])
				logger := logrus.WithFields(logrus.Fields{
					"request": string(httpRequest),
					"stack":   string(buf[:n]),
				})
				requestID := c.GetString(constant.XRequestID)
				if requestID != "" {
					logger = logger.WithField("request_id", requestID)
				}
				logger.Error(fmt.Sprintf("%#v", err))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}

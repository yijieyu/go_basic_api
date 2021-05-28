package middleware

import (
	"bytes"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thinkeridea/go-extend/exbytes"
	"github.com/thinkeridea/go-extend/exstrings"

	"github.com/yijieyu/go_basic_api/pkg/constant"
	"github.com/yijieyu/go_basic_api/pkg/logger"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *bodyLogWriter) WriteString(s string) (n int, err error) {
	return w.Write(exstrings.UnsafeToBytes(s))
}

// Logging is a middleware function that logs the each request.
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logrus.WithFields(logrus.Fields{
			"request_id": c.GetString(constant.XRequestID),
		})

		c.Set(constant.Logger, log)

		if !logger.Debug() {
			return
		}

		start := time.Now().UTC()

		// Read the Body content
		var bodyBytes []byte
		if c.Request.Body != nil && !strings.Contains(c.Request.Header.Get("Content-type"), "multipart/form-data") {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		log = log.WithFields(logrus.Fields{
			"query":        c.Request.URL.String(),
			"method":       c.Request.Method,
			"content_type": c.GetHeader("Content-type"),
			"client_ip":    c.GetString(constant.HttpClientIP),
		})

		log.WithField("body", exbytes.ToString(bodyBytes)).Debug("request")

		// Restore the io.ReadCloser to its original state
		blw := &bodyLogWriter{
			body:           bytes.NewBuffer(make([]byte, 0, 1024)),
			ResponseWriter: c.Writer,
		}
		c.Writer = blw

		// Continue.
		c.Next()

		log.WithFields(logrus.Fields{
			"body":     exbytes.ToString(blw.body.Bytes()),
			"duration": time.Now().UTC().Sub(start).Milliseconds(),
		}).Debug("response")
	}
}

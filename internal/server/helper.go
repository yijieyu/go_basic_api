package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rs/xid"
	"github.com/sirupsen/logrus"
	"github.com/thinkeridea/go-extend/exnet"

	"github.com/yijieyu/go_basic_api/internal/apperr"
	"github.com/yijieyu/go_basic_api/pkg/constant"
	"github.com/yijieyu/go_basic_api/pkg/errno"
)

func init() {
	binding.Validator = nil
}

type Validate interface {
	Validate() error
}

func Bind(c *gin.Context, v interface{}, binding ...binding.Binding) error {
	var err error
	if len(binding) <= 0 {
		err = c.ShouldBind(v)
		if err != nil {
			return apperr.Param.Wrap(err)
		}
	}

	for _, bind := range binding {
		err = c.ShouldBindWith(v, bind)
		if err != nil {
			return apperr.Param.Wrap(err)
		}
	}

	if validate, ok := v.(Validate); ok {
		if err = validate.Validate(); err != nil {
			return apperr.Validation.WrapComment(err)
		}
	}

	return nil
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(err error, data ...interface{}) *Response {
	res := &Response{}
	for i := 0; i < len(data); i++ {
		res.Data = data[i]
	}
	res.Code, res.Message, res.Error = errno.Decode(err)
	return res
}

func RequestID(c *gin.Context) string {
	requestID := c.GetString(constant.XRequestID)
	if requestID != "" {
		return requestID
	}

	requestID = c.Request.Header.Get(constant.XRequestID)
	if requestID == "" {
		requestID = xid.New().String()
	}

	c.Set(constant.XRequestID, requestID)
	c.Writer.Header().Set(constant.XRequestID, requestID)
	return requestID
}

func Logger(c *gin.Context) *logrus.Entry {
	var logger *logrus.Entry
	v, ok := c.Get(constant.Logger)
	if !ok {
		goto RETURN
	}

	logger, ok = v.(*logrus.Entry)
	if ok {
		return logger
	}

RETURN:
	logger = logrus.WithFields(logrus.Fields{
		"request_id": RequestID(c),
	})
	c.Set(constant.Logger, logger)
	return logger
}

func ClientIP(c *gin.Context) string {
	ip := c.GetString(constant.HttpClientIP)
	if ip != "" {
		return ip
	}

	ip = exnet.ClientPublicIP(c.Request)
	if ip == "" {
		ip = exnet.ClientIP(c.Request)
	}

	c.Set(constant.HttpClientIP, ip)
	return ip
}

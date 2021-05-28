package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yijieyu/go_basic_api/pkg/logger"
)

const logLevelAutoRecoverInterval = 30 * time.Minute

type Assist struct {
}

func NewAssist() *Assist {
	return &Assist{}
}

func (h *Assist) SetLogLevel(defaultLogLevel string) gin.HandlerFunc {
	return func(c *gin.Context) {
		level := c.Param("level")
		currentLevel := logrus.GetLevel().String()
		err := logger.SetLogLevel(level)
		if err != nil {
			c.String(http.StatusOK, err.Error())
			return
		}

		time.AfterFunc(logLevelAutoRecoverInterval, func() {
			log := Logger(c).WithFields(logrus.Fields{
				"current_level": logrus.GetLevel().String(),
				"default_level": defaultLogLevel,
			})

			if err := logger.SetLogLevel(defaultLogLevel); err != nil {
				log.Error(err.Error())
				return
			}

			log.Infof("the log level recover to default %s", defaultLogLevel)
		})

		c.String(http.StatusOK, fmt.Sprintf("Modify %s level to %s", currentLevel, level))
	}
}

func (h *Assist) Reload(reload func(...string) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		var events []string
		if es := c.Query("events"); len(es) > 0 {
			events = strings.Split(es, ",")
		}

		log := Logger(c)
		if err := reload(events...); err != nil {
			log.WithFields(logrus.Fields{
				"events": events,
			}).Error(err.Error())

			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		log.WithFields(logrus.Fields{
			"events": events,
		}).Info("reload succeed")
		c.String(http.StatusOK, "ok")
	}
}

func (h *Assist) Debug(f func(cmd string, c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd := c.Param("cmd")
		if err := f(cmd, c); err != nil {
			log := Logger(c).WithFields(logrus.Fields{
				"command": cmd,
			})

			log.Error(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
}

package logger

import (
	"github.com/sirupsen/logrus"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/configuration"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func InitSqlLog(conf configuration.SqlLog) *logrus.Entry {

	logger := logrus.New()

	l, err := logrus.ParseLevel(conf.Level)
	if err == nil {
		logger.SetLevel(l)
	}

	rotate := &lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    500, // 最大的文件500M
		MaxBackups: 10,  // 最多保留10个文件
		MaxAge:     7,   // 最长保留7天
	}

	if conf.MaxAge > 0 {
		rotate.MaxAge = conf.MaxAge
	}

	if conf.MaxSize > 0 {
		rotate.MaxSize = conf.MaxSize
	}

	switch conf.Format {
	case logTypeText:
		logger.SetFormatter(&TextFormatter{TextFormatter: logrus.TextFormatter{DisableColors: true}})
	case logTypeJSON:
		logger.SetFormatter(&JSONFormatter{JSONFormatter: logrus.JSONFormatter{}})
	default:
	}

	logger.SetOutput(rotate)
	entry := logrus.NewEntry(logger)

	return entry
}

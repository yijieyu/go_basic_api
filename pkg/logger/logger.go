package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/yijieyu/go_basic_api/pkg/configuration"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

const (
	logTypeText = "text"
	logTypeJSON = "json"
)

func InitLog(conf configuration.Log) error {
	var err error

	err = SetLogLevel(conf.Level)
	if err != nil {
		return err
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

	if conf.MaxBackups > 0 {
		rotate.MaxBackups = conf.MaxBackups
	}

	if conf.LocalTime {
		rotate.LocalTime = conf.LocalTime
	}

	if conf.Compress {
		rotate.Compress = conf.Compress
	}

	logrus.SetOutput(rotate)

	// 使用无锁模式
	logrus.StandardLogger().SetNoLock()

	switch conf.Format {
	case logTypeText:
		logrus.SetFormatter(&TextFormatter{TextFormatter: logrus.TextFormatter{DisableColors: true}, EnableFuncCallDepthRecord: conf.Stack})
	case logTypeJSON:
		fallthrough
	default:
		logrus.SetFormatter(&JSONFormatter{EnableFuncCallDepthRecord: conf.Stack})
	}

	return nil
}

func SetLogLevel(level string) error {
	l, err := logrus.ParseLevel(level)
	if err == nil {
		logrus.SetLevel(l)
	}

	return err
}

func Debug() bool {
	return logrus.GetLevel() == logrus.DebugLevel
}

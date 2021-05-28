package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

func NewRecordLog(w io.Writer) *logrus.Entry {
	logger := &logrus.Logger{
		Out:       w,
		Formatter: new(logrus.JSONFormatter),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}

	return logrus.NewEntry(logger)
}

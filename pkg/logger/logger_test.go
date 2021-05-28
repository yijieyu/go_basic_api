package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestWithFields(t *testing.T) {
	Convey("Given Logger WithFields", t, func() {
		var buffer bytes.Buffer
		var fields logrus.Fields
		logrus.SetOutput(&buffer)

		logger := logrus.WithFields(logrus.Fields{"test": "test"})

		Convey("The when Json format log test", func() {
			logrus.SetFormatter(&JSONFormatter{EnableFuncCallDepthRecord: true})
			logger.Info("info test")
			_, file, line, _ := runtime.Caller(0)

			Convey("The logger fields is equal to the expected", func() {
				err := json.Unmarshal(buffer.Bytes(), &fields)
				So(err, ShouldBeNil)

				So(fields["msg"], ShouldEqual, "info test")
				So(fields["test"], ShouldEqual, "test")
				So(fields["level"], ShouldEqual, logrus.InfoLevel.String())
				So(fields["stack"], ShouldEqual, fmt.Sprintf("%s:%d", file, line-1))
			})
		})

		Convey("The when String format log test", func() {
			logrus.SetFormatter(&TextFormatter{
				TextFormatter:             logrus.TextFormatter{DisableColors: true},
				EnableFuncCallDepthRecord: true,
			})

			logger.Info("info_test")
			_, file, line, _ := runtime.Caller(0)

			Convey("The logger fields is equal to the expected", func() {
				fields = make(logrus.Fields)
				for _, kv := range strings.Split(strings.Trim(buffer.String(), "\n"), " ") {
					if !strings.Contains(kv, "=") {
						continue
					}
					kvArr := strings.Split(kv, "=")
					key := strings.TrimSpace(kvArr[0])
					val := kvArr[1]
					if kvArr[1][0] == '"' {
						var err error
						val, err = strconv.Unquote(val)
						So(err, ShouldBeNil)
					}
					fields[key] = val
				}

				So(fields["msg"], ShouldEqual, "info_test")
				So(fields["test"], ShouldEqual, "test")
				So(fields["level"], ShouldEqual, logrus.InfoLevel.String())
				So(fields["stack"], ShouldEqual, fmt.Sprintf("%s:%d", file, line-1))
			})
		})
	})
}

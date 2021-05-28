package api

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAssistAPIHandler_SetLogLevel(t *testing.T) {
	Convey("Given SetLogLevel handlers", t, func() {
		handler := NewAssist()
		resp := httptest.NewRecorder()
		_, r := gin.CreateTestContext(resp)
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.ErrorLevel)

		r.GET("/log/:level", handler.SetLogLevel(logrus.ErrorLevel.String()))

		Convey("The correct request parameters", func() {
			req, _ := http.NewRequest("GET", "/log/debug", nil)
			r.ServeHTTP(resp, req)

			Convey("The current level of error is not equal to the debug", func() {
				So(resp.Code, ShouldEqual, http.StatusOK)
				So(logrus.GetLevel(), ShouldEqual, logrus.DebugLevel)
			})
		})

		Convey("Set does not exist log level", func() {
			req, _ := http.NewRequest("GET", "/log/unknown", nil)
			r.ServeHTTP(resp, req)

			Convey("The current level of error is not equal to the unknown", func() {
				So(resp.Code, ShouldEqual, http.StatusOK)
				So(logrus.GetLevel().String(), ShouldNotEqual, "unknown")
			})
		})
	})
}

type reloadMock struct {
	events []string
}

func (m *reloadMock) OnReload(event ...string) error {
	m.events = event
	if len(event) > 0 && event[0] == "error" {
		return errors.New("reload fail")
	}
	return nil
}

func TestAssistAPIHandler_Reload(t *testing.T) {
	Convey("reload cache api test", t, func() {
		handler := NewAssist()
		resp := httptest.NewRecorder()
		_, r := gin.CreateTestContext(resp)

		c := &reloadMock{}
		r.GET("/reload", handler.Reload(c.OnReload))

		Convey("reload all cache test", func() {
			req, _ := http.NewRequest("GET", "/reload", nil)
			r.ServeHTTP(resp, req)

			Convey("Reload the success, event list is empty", func() {
				So(resp.Code, ShouldEqual, http.StatusOK)
				So(len(c.events), ShouldEqual, 0)

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Errorf("failed in ReadAll resp.Body:%v", err)
				}

				So(body, ShouldResemble, []byte("ok"))
			})
		})

		Convey("reload events cache test", func() {
			req, _ := http.NewRequest("GET", "/reload?events=a,b,c,d,e,f", nil)
			r.ServeHTTP(resp, req)

			Convey("Reload the success, event list and expectations", func() {
				So(resp.Code, ShouldEqual, http.StatusOK)
				So(c.events, ShouldResemble, []string{"a", "b", "c", "d", "e", "f"})

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Errorf("failed in ReadAll resp.Body:%v", err)
				}

				So(body, ShouldResemble, []byte("ok"))
			})
		})

		Convey("reload failed test", func() {
			req, _ := http.NewRequest("GET", "/reload?events=error,b,c,d,e,f", nil)
			r.ServeHTTP(resp, req)

			Convey("Reload the failed", func() {
				So(resp.Code, ShouldEqual, http.StatusInternalServerError)

				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Errorf("failed in ReadAll resp.Body:%v", err)
				}

				So(body, ShouldNotResemble, []byte("ok"))
			})
		})
	})
}

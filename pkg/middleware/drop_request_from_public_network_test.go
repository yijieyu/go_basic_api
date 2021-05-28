package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDropRequestFromPublicNetwork(t *testing.T) {
	Convey("Given drop request from public network middleware", t, func() {
		Convey("When request from public network", func() {
			resp := httptest.NewRecorder()
			_, r := gin.CreateTestContext(resp)

			route := "/route"
			r.GET(route, DropRequestFromPublicNetwork, func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			for _, v := range []struct {
				xForwardedFor string
				remoteAddr    string
				expected      int
			}{
				{"10.3.5.45", "10.10.25.176:100", http.StatusOK},
				{"", "127.0.0.1:100", http.StatusOK},
				{"", "192.168.9.18:100", http.StatusOK},
				{"", "10.168.9.18:100", http.StatusOK},
				{"", "172.16.9.18:100", http.StatusOK},

				{"10.3.5.45, 21.45.9.1", "101.1.0.4:100", http.StatusForbidden},
				{"101.3.5.45, 21.45.9.1", "101.1.0.4:100", http.StatusForbidden},
				{"", "101.1.0.4:100", http.StatusForbidden},
				{"21.45.9.1", "101.1.0.4:100", http.StatusForbidden},
				{"21.45.9.1, ", "101.1.0.4:100", http.StatusForbidden},
				{"192.168.5.45, 210.45.9.1, 89.5.6.1", "101.1.0.4:100", http.StatusForbidden},
				{"192.168.5.45, 172.24.9.1, 89.5.6.1", "101.1.0.4:100", http.StatusForbidden},
				{"192.168.5.45, 172.24.9.1", "101.1.0.4:100", http.StatusForbidden},
				{"192.168.5.45, 172.24.9.1", "101.1.0.4:5670", http.StatusForbidden},
			} {
				Convey(fmt.Sprintf("IsxForwardedFor:%s, remoteAddr:%s, http code Should Equal %d", v.xForwardedFor, v.remoteAddr, v.expected), func() {
					req, _ := http.NewRequest("GET", route, nil)
					req.Header.Add("X-Forwarded-For", v.xForwardedFor)
					req.RemoteAddr = v.remoteAddr
					r.ServeHTTP(resp, req)

					So(resp.Code, ShouldEqual, v.expected)
				})
			}
		})
	})
}

package ad

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sirupsen/logrus"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/model"
	api "gitlab.weimiaocaishang.com/weimiao/base_api/internal/server"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/service"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/configuration"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/errno"
)

type Ad struct {
	s             *service.Ad
	thirdAdSwitch configuration.ThirdAdSwitch
}

func NewAd(s *service.Ad, thirdAdSwitch configuration.ThirdAdSwitch) *Ad {
	return &Ad{
		s:             s,
		thirdAdSwitch: thirdAdSwitch,
	}
}

func (a *Ad) GetAd(c *gin.Context) {

	var req model.Request
	var err error
	var res model.ResponseAd

	logger := api.Logger(c)

	defer func() {
		c.JSON(http.StatusOK, api.NewResponse(err, res))
	}()

	err = api.Bind(c, &req, binding.JSON)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error("get ad request param err")
		err = errno.New(70001, "request ad param error")
		return
	}

	if req.App == nil || req.Device == nil || req.Ad == nil {
		logger.WithFields(logrus.Fields{
			"req": req,
		}).Error("get ad request param err")
		err = errno.New(70001, "request ad param error")
		return
	}

	req.Device.IP = api.ClientIP(c)
	req.RequestID = api.RequestID(c)

	if req.Device.IP != "" {
		fmt.Println(req.Device.IP)
	}

	req.User = &model.RequestUser{
		ID: 0,
	}

	res = a.s.GetAdPlaceID(req, logger)

}

func (a *Ad) GetThirdAdSwitch(c *gin.Context) {

	var resp model.ThirdAdSwitchResp
	if a.thirdAdSwitch.Banner == "1" {
		resp.Banner = true
	}
	if a.thirdAdSwitch.OpenScreen == "1" {
		resp.OpenScreen = true
	}
	c.JSON(http.StatusOK, api.NewResponse(nil, resp))

}

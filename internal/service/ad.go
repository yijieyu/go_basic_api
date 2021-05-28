package service

import (
	"github.com/go-playground/form"
	"github.com/sirupsen/logrus"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/model"
)

type Ad struct {
	reportURL   string
	formEncoder *form.Encoder
}

func NewAd(reportURL string) *Ad {
	return &Ad{
		reportURL:   reportURL,
		formEncoder: form.NewEncoder(),
	}
}

func (s *Ad) GetAdPlaceID(req model.Request, logger *logrus.Entry) model.ResponseAd {

	return model.ResponseAd{}
}

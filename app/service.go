package app

import "gitlab.weimiaocaishang.com/weimiao/base_api/internal/service"

type Service struct {
	ad            *service.Ad
	infoflowMedia *service.InfoflowMediaService
}

func (s *Service) init(app *App) {
	s.ad = service.NewAd(app.config.ReportURL)
	s.infoflowMedia = service.NewInfoflowMediaService(app.Storage.mysql)
}

func (s *Service) InfoflowMedia() *service.InfoflowMediaService {
	return s.infoflowMedia
}
func (s *Service) Ad() *service.Ad {
	return s.ad
}

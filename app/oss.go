package app

import (
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/provider/oss"
)

type Oss struct {
	oss    *oss.AliyunOSS
	ResURL string
	Tmp    string
}

func (o *Oss) init(app *App) {

}

func (o *Oss) Oss() *oss.AliyunOSS {
	return o.oss
}

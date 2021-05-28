package app

import (
	"bytes"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/thinkeridea/go-extend/helper"
	"gitlab.weimiaocaishang.com/weimiao/base_api/internal/model"
	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/provider/apollo"

	"gitlab.weimiaocaishang.com/weimiao/base_api/pkg/configuration"
)

const (
	ConfEnv = "API"
)

// Configuration model
type Config struct {
	Build   configuration.Build   `mapstructure:",squash" validate:"required"`
	Service configuration.Service `mapstructure:"service" validate:"required"`
	HTTP    configuration.HTTP    `mapstructure:"http" validate:"required"`
	Log     configuration.Log     `mapstructure:"log" validate:"required"`

	DB struct {
		Master []configuration.DB `mapstructure:"master" validate:"required"`
		Slave  []configuration.DB `mapstructure:"slave"`
	} `mapstructure:"db" validate:"required"`

	Redis struct {
		Master []configuration.Redis `mapstructure:"master" validate:"required"`
		Slave  []configuration.Redis `mapstructure:"slave"`
	} `mapstructure:"redis" validate:"required"`

	Etcd            []string `mapstructure:"etcd" validate:"required"`
	ConfCachePrefix string   `mapstructure:"conf_cache_prefix" validate:"required"`

	ReportURL string `mapstructure:"report_url" validate:"required"`

	configuration.Jwt `mapstructure:"jwt" validate:"required"`

	configuration.ThirdAdSwitch `mapstructure:"third_ad_switch" validate:"required"`
}

func InitConfig(app *App) {

	config := new(Config)

	viper.SetEnvPrefix(ConfEnv)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	client := apollo.NewApolloClient(model.AppID, app.env)

	confByte := helper.Must(client.GetNoCacheConfig()).([]byte)

	helper.Must(nil, viper.ReadConfig(bytes.NewBuffer(confByte)))
	helper.Must(nil, viper.Unmarshal(config))
	helper.Must(nil, configuration.Validate(config))

	app.rw.Lock()
	app.config = config
	app.rw.Unlock()

	go func() {
		for {
			select {
			default:
				// check 配置信息是否有变化
				isChange := client.V2Request()

				// 配置有变化，获取最新的配置
				if !isChange {
					continue
				}
				value, err := client.GetCacheConfig()
				if err != nil {
					continue
				}
				confRes := ReloadConfig(value)
				if confRes != nil {
					app.rw.Lock()
					app.config = confRes
					app.rw.Unlock()
				}
			}
		}
	}()
}

func ReloadConfig(confByte []byte) *Config {

	config := new(Config)

	err := viper.ReadConfig(bytes.NewBuffer(confByte))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"resp_byte": string(confByte),
			"err":       err,
		}).Error("app.viper.ReadConfig err")
		return nil
	}
	err = viper.Unmarshal(config)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"resp_byte": string(confByte),
			"err":       err,
		}).Error("app.viper.Unmarshal err")
		return nil
	}
	err = configuration.Validate(config)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"resp_byte": string(confByte),
			"err":       err,
		}).Error("configuration.Validate err")
		return nil
	}

	return config

}

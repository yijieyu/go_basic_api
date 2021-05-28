package configuration

type ThirdAdSwitch struct {
	Banner     string `mapstructure:"banner" validate:"required"`
	OpenScreen string `mapstructure:"open_screen" validate:"required"`
}

package configuration

type Sms struct {
	URL    string    `mapstructure:"url" validate:"required"`
	Sign   string    `mapstructure:"sign" validate:"required"`
	Notice SmsNotice `mapstructure:"notice" validate:"required"`
	Market SmsMarket `mapstructure:"market" validate:"required"`
}

type SmsNotice struct {
	Account  string `mapstructure:"account" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}
type SmsMarket struct {
	Account  string `mapstructure:"account" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
}

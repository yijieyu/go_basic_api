package configuration

type Jwt struct {
	SecretKey string `mapstructure:"secret_key" validate:"required"`
}

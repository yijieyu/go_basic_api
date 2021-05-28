package configuration

type Oss struct {
	AccessKeyID     string `mapstructure:"access_key_id" validate:"required"`
	AccessKeySecret string `mapstructure:"access_key_secret" validate:"required"`
	Endpoint        string `mapstructure:"endpoint" validate:"required"`
	BucketName      string `mapstructure:"bucket_name" validate:"required"`
	ResURL          string `mapstructure:"res_url" validate:"required"`
	Tmp             string `mapstructure:"tmp" validate:"required"`
}

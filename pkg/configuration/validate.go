package configuration

import (
	"time"

	validator "github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
)

// Validate validates configuration
func Validate(c interface{}) error {
	validate := validator.New()
	validate.SetTagName("validate")
	if err := validate.RegisterValidation("dsn", dsnValidator); err != nil {
		return err
	}

	if err := validate.RegisterValidation("duration", durationValidator); err != nil {
		return err
	}

	return validate.Struct(c)
}

func dsnValidator(fl validator.FieldLevel) bool {
	dsn := fl.Field().String()
	_, err := mysql.ParseDSN(dsn)
	return err == nil
}

func durationValidator(fl validator.FieldLevel) bool {
	duration := fl.Field().String()
	_, err := time.ParseDuration(duration)
	return err == nil
}

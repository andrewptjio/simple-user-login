package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// validator
type UserValidator struct {
	Validator *validator.Validate
}

// validator validation
func (v UserValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, err.Error())
	}

	return nil
}

package user

import (
	"net/http"
	"simple-user-login/model"
	"simple-user-login/repository"

	"github.com/labstack/echo/v4"
)

// request
type requestCreate struct {
	UserName string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Role     string `json:"role" form:"role" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Age      int    `json:"age" form:"age" validate:"required"`
}

// response
type responseCreate struct {
	User  *model.User `json:"user"`
	Token string      `json:"token"`
}

// business logic validation
func (r *requestCreate) Check() error {
	var err error

	// check if duplicate username
	if _, err = repository.GetUser(r.UserName); err == nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Username existed")
	}

	// check role, can only be admin or user
	if r.Role != "admin" && r.Role != "user" {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Invalid Role")
	}

	// check password length, must at least 8 characters
	if len(r.Password) < 8 {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Password has to be more than 8 characters")
	}

	// check age if above 17 and below 100
	if r.Age < 17 {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Age must be greater than 17")
	}
	if r.Age > 100 {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Age must be less than 100")
	}

	return nil
}

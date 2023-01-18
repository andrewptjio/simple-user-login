package user

import (
	"net/http"
	"simple-user-login/model"
	"simple-user-login/repository"

	"github.com/labstack/echo/v4"
)

// request
type requestDelete struct {
	UserName string `json:"username" form:"username" validate:"required"`

	User *model.User
}

// response
type responseDelete struct {
	User *model.User `json:"user"`
}

// business logic validation
func (r *requestDelete) Check() error {
	var err error

	// check user
	if r.User, err = repository.GetUser(r.UserName); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Wrong Username/Password")
	}

	return nil
}

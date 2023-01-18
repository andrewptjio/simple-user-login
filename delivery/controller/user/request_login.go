package user

import (
	"net/http"
	"simple-user-login/model"
	"simple-user-login/repository"
	"simple-user-login/util"

	"github.com/labstack/echo/v4"
)

// request
type requestLogin struct {
	UserName string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`

	User *model.User
}

// response
type responseLogin struct {
	Token string `json:"token"`
}

// business logic validation
func (r *requestLogin) Check() error {
	var err error

	// check user
	if r.User, err = repository.GetUser(r.UserName); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Wrong Username/Password")
	}

	// check password
	if _, err = util.Checkpwd(r.User.Password, r.Password); err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "Wrong Username/Password")
	}

	return nil
}

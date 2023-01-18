package user

import (
	"fmt"
	"net/http"
	"simple-user-login/delivery/middleware"
	"simple-user-login/model"
	"simple-user-login/repository"

	"github.com/labstack/echo/v4"
)

// request
type requestUpdate struct {
	Name string `json:"name" form:"name" validate:"required"`
	Age  int    `json:"age" form:"age" validate:"required"`

	// admin edit
	Username string `json:"username" form:"username"`
	Role     string `json:"role" form:"role"`

	User    *model.User
	Session middleware.JWTPayload
}

// response
type responseUpdate struct {
	User *model.User `json:"user"`
}

// business logic validation
func (r *requestUpdate) Check() error {
	var err error

	// check if admin is trying to update someone account
	if r.Session.Role == "admin" && r.Session.Username != r.Username {
		// check user
		if r.User, err = repository.GetUser(r.Username); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Username doesn't exist")
		}

		// check role, can only be admin or user
		if r.Role != "admin" && r.Role != "user" {
			return echo.NewHTTPError(http.StatusNotAcceptable, "Invalid Role")
		}
	} else {
		// user update
		if r.User, err = repository.GetUser(r.Session.Username); err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Wrong Username/Password")
		}
		r.Role = "user"
	}

	// check age if above 17 and below 100
	if r.Age < 17 {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Age must be greater than 17")
	}
	if r.Age > 100 {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Age must be less than 100")
	}

	fmt.Println("usernya ", r.User)

	return nil
}

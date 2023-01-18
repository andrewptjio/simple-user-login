package user

import (
	"net/http"
	"simple-user-login/config"
	mw "simple-user-login/delivery/middleware"
	"simple-user-login/model"
	"simple-user-login/repository"
	"simple-user-login/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRoute(e *echo.Echo) {
	e.POST("/user/register", registerUser())
	e.POST("/user/login", loginUser())

	e.GET("/user", readAllUser(), middleware.JWT([]byte(config.JWTSecretKey)))
	e.GET("/user/profile", profileUser(), middleware.JWT([]byte(config.JWTSecretKey)))
	e.PUT("/user", editUser(), middleware.JWT([]byte(config.JWTSecretKey)))
	e.DELETE("/user", removeUser(), middleware.JWT([]byte(config.JWTSecretKey)))
}

// registerUser : function to handle user registration
func registerUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			r    *requestCreate
			data *responseCreate
		)
		c.Bind(&r)

		// validator
		if err := c.Validate(r); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.NewStatusNotAcceptable())
		}

		// check logic
		if err := r.Check(); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.ErrorResponse(400, err))
		}

		// create user
		user, err := save(r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, util.ErrorResponse(500, err))
		}

		// create token
		token, _ := mw.CreateToken(user.ID, user.UserName, user.Role)

		data = &responseCreate{
			User:  user,
			Token: token,
		}

		return c.JSON(http.StatusOK, util.SuccessResponse(data))
	}
}

// loginUser : function to handle login
func loginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			r    *requestLogin
			data *responseLogin
		)
		c.Bind(&r)

		// validator
		if err := c.Validate(r); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.NewStatusNotAcceptable())
		}

		// check logic
		if err := r.Check(); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.ErrorResponse(400, err))
		}

		// create token
		token, _ := mw.CreateToken(r.User.ID, r.User.UserName, r.User.Role)

		data = &responseLogin{
			Token: token,
		}

		return c.JSON(http.StatusOK, util.SuccessResponse(data))
	}
}

// readAllUser : function to read all data
func readAllUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data []*model.User
		var err error

		// check session
		session, _ := mw.ExtractTokenUser(c)
		if session.Role != "admin" {
			return c.JSON(http.StatusForbidden, util.NewStatusForbidden())
		}

		// get data
		if data, err = repository.GetAllUser(); err != nil {
			return c.JSON(http.StatusInternalServerError, util.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, util.SuccessResponse(data))
	}
}

// profileUser : function to read user's profile
func profileUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var data *model.User
		var err error

		// only admin can access this feature
		session, _ := mw.ExtractTokenUser(c)

		// get data
		if data, err = repository.GetUser(session.Username); err != nil {
			return c.JSON(http.StatusInternalServerError, util.NewInternalServerErrorResponse())
		}

		return c.JSON(http.StatusOK, util.SuccessResponse(data))
	}
}

// editUser : function to edit user
func editUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			r    *requestUpdate
			data *responseUpdate
		)
		c.Bind(&r)

		// check if user has token
		r.Session, _ = mw.ExtractTokenUser(c)

		// validator
		if err := c.Validate(r); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.NewStatusNotAcceptable())
		}

		// check logic
		if err := r.Check(); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.ErrorResponse(400, err))
		}

		// update user
		user, err := update(r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, util.ErrorResponse(500, err))
		}

		data = &responseUpdate{
			User: user,
		}

		return c.JSON(http.StatusOK, util.SuccessResponse(data))
	}
}

// removeUser : function to remove user
func removeUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			r    *requestDelete
			data *responseDelete
		)

		c.Bind(&r)

		// check if user has token
		session, _ := mw.ExtractTokenUser(c)
		if session.Role != "admin" {
			return c.JSON(http.StatusForbidden, util.NewStatusForbidden())
		}

		// validator
		if err := c.Validate(r); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.NewStatusNotAcceptable())
		}

		// check logic
		if err := r.Check(); err != nil {
			return c.JSON(http.StatusNotAcceptable, util.ErrorResponse(400, err))
		}

		// delete user
		user, err := delete(r)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, util.ErrorResponse(500, err))
		}

		data = &responseDelete{
			User: user,
		}

		return c.JSON(http.StatusOK, util.SuccessResponse(data))
	}
}

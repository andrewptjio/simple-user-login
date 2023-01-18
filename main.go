package main

import (
	"simple-user-login/config"

	"simple-user-login/delivery/controller/user"
	mw "simple-user-login/delivery/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// init config
	config.InitConfig()

	// mongo connection tester
	config.MongoConnectionTest()

	e := echo.New()

	mw.LogMiddleware(e)

	e.Pre(middleware.RemoveTrailingSlash())

	// validator
	e.Validator = &user.UserValidator{Validator: validator.New()}

	// register route
	user.UserRoute(e)

	e.Logger.Fatal(e.Start(":" + config.AppPort))
}

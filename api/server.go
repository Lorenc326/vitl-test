package api

import (
	"github.com/Lorenc326/vitl-test/number"
	"github.com/Lorenc326/vitl-test/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BuildServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/generate", number.Generate)

	e.POST("/register", user.Register)
	e.POST("/login", user.Login)

	e.GET("/details", user.Details, Authorized)

	return e
}

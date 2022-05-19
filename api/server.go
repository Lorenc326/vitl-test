package api

import (
	"github.com/Lorenc326/vitl-test/number"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BuildServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/generate", number.Generate)

	return e
}

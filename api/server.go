package api

import (
"github.com/labstack/echo/v4"
"github.com/labstack/echo/v4/middleware"
)

func BuildServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}
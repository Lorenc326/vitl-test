package api

import (
	"github.com/Lorenc326/vitl-test/services/auth"
	"github.com/labstack/echo/v4"
)

func Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(auth.CookieName)
		if err == nil {
			email, err := auth.Authenticate(cookie)
			if err == nil {
				c.Set("email", email)
			}
		}
		return next(c)
	}
}

func Protected(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(auth.CookieName)
		if err != nil {
			return echo.ErrUnauthorized
		}
		email, err := auth.Authenticate(cookie)
		if err != nil {
			return echo.ErrUnauthorized
		}
		c.Set("email", email)
		return next(c)
	}
}

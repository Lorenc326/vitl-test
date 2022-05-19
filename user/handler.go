package user

import (
	"github.com/Lorenc326/vitl-test/services/auth"
	"github.com/Lorenc326/vitl-test/services/validator"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type RegisterInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Name     string `json:"name" validate:"required,min=4,max=15"`
}

func Register(c echo.Context) error {
	input := RegisterInput{}
	if err := c.Bind(&input); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	if err := validator.Validate.Struct(input); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := createUser(input); err != nil {
		return echo.ErrInternalServerError
	}

	return nil
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func Login(c echo.Context) error {
	input := LoginInput{}
	if err := c.Bind(&input); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	if err := validator.Validate.Struct(input); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := authenticate(input.Email, input.Password); err != nil {
		return echo.ErrBadRequest
	}
	c.SetCookie(auth.GetAuthCookie(input.Email))

	return nil
}

func Details(c echo.Context) error {
	email := c.Get("email") // preset by middleware
	user := getUser(email.(string))
	user.Email = "" // passport is not exported
	c.JSON(200, user)
	return nil
}

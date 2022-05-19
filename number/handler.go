package number

import (
	"github.com/Lorenc326/vitl-test/services/validator"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type GenerateInput struct {
	FromNumber   int `json:"from_number" validate:"required"`
	ToNumber     int `json:"to_number" validate:"required"`
	TotalNumbers int `json:"total_numbers" validate:"required"`
}

func Generate(c echo.Context) error {
	sort := c.QueryParam("sort")
	if sort != "" && sort != "asc" && sort != "desc" {
		return echo.ErrBadRequest
	}
	if sort == "" {
		sort = "asc"
	}

	input := GenerateInput{}
	if err := c.Bind(&input); err != nil {
		log.Println(err)
		return echo.ErrBadRequest
	}
	if err := validator.Validate.Struct(input); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	counts, err := countsForRandomSequence(sort, &input)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(200, counts)
}

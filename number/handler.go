package number

import (
	"github.com/labstack/echo/v4"
	"log"
)

type GenerateInput struct {
	FromNumber   int `json:"from_number"`
	ToNumber     int `json:"to_number"`
	TotalNumbers int `json:"total_numbers"`
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

	counts, err := countsForRandomSequence(sort, &input)
	if err != nil {
		log.Println(err)
		return echo.ErrInternalServerError
	}
	return c.JSON(200, counts)
}

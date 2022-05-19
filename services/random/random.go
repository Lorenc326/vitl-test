package random

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Config struct {
	Min    int
	Max    int
	Amount int
}

const basePath = "https://www.random.org/integers"

func GetSequenceRandomNumber(c Config) ([]int, error) {
	url := fmt.Sprintf("%s/?num=%d&min=%d&max=%d&col=1&base=10&format=plain", basePath, c.Amount, c.Min, c.Max)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	return textToNumbers(string(body)), err
}

func textToNumbers(sequence string) []int {
	// => "10 3 1 3"
	trimmed := strings.Trim(sequence, "\n")

	// => []string{"10", "3", "1", "3"}
	strings := strings.Split(trimmed, "\n")

	res := make([]int, len(strings))
	for i, s := range strings {
		res[i], _ = strconv.Atoi(s)
	}

	return res
}
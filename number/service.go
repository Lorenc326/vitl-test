package number

import (
	"github.com/Lorenc326/vitl-test/services/random"
)

func countsForRandomSequence(order string, input *GenerateInput) (*[]int, error) {
	randomSequence, err := random.GetSequenceRandomNumber(random.Config{
		Min:    input.FromNumber,
		Max:    input.ToNumber,
		Amount: input.TotalNumbers,
	})
	if err != nil {
		return nil, err
	}
	return countDuplicates(order, randomSequence, input.FromNumber, input.ToNumber), nil
}

func countDuplicates(order string, set []int, min, max int) *[]int {
	counts := make([]int, max-min+1)

	if order == "asc" {
		for _, num := range set {
			counts[num-min]++
		}
	} else {
		for _, num := range set {
			counts[max-num]++
		}
	}

	return &counts
}

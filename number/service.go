package number

import (
	"github.com/Lorenc326/vitl-test/services/random"
	"sort"
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
	return countDuplicates(order, randomSequence), nil
}

func countDuplicates(order string, set []int) *[]int {
	counts := make(map[int]int)
	uniqueNumbers := make([]int, 0)
	for _, num := range set {
		if counts[num] == 0 {
			uniqueNumbers = append(uniqueNumbers, num)
		}
		counts[num]++
	}

	if order == "asc" {
		sort.Sort(sort.IntSlice(uniqueNumbers))
	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(uniqueNumbers)))
	}

	// replace number with its count
	for i := 0; i < len(uniqueNumbers); i++ {
		uniqueNumbers[i] = counts[uniqueNumbers[i]]
	}

	return &uniqueNumbers
}

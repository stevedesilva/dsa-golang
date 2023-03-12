package applications

import (
	"errors"

	"github.com/stevedesilva/dsa-golang.git/commonsense/sort"
)

var (
	MinimumErr      = errors.New("minimum args not met")
	ItemNotFoundErr = errors.New("minimum args not met")
)

func FindMissingNumber(numbers []int) (*int, error) {
	if numbers == nil || len(numbers) < 1 {
		return nil, MinimumErr
	}
	sort.NewQuickSort(numbers).Quicksort()
	for i := 0; i < len(numbers); i++ {
		if numbers[i] != i {
			return IntPointer(i), nil
		}
	}

	return nil, ItemNotFoundErr
}

func IntPointer(value int) *int {
	return &value
}

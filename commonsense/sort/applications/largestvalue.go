package applications

import (
	"errors"

	"github.com/stevedesilva/dsa-golang.git/commonsense/sort"
)

var (
	ErrMinimumValue = errors.New("minimum values of slice not supplied")
)

func LargestValueN(numbers []int) (*int, error) {
	if len(numbers) < 1 {
		return nil, ErrMinimumValue
	}
	maxValue := numbers[0]
	for i := 1; i < len(numbers); i++ {
		if maxValue < numbers[i] {
			maxValue = numbers[i]
		}
	}
	return intPtr(maxValue), nil
}

func LargestValueNLogN(numbers []int) (*int, error) {
	if len(numbers) < 1 {
		return nil, ErrMinimumValue
	}
	sort.NewQuickSort(numbers).Quicksort()
	return intPtr(numbers[len(numbers)-1]), nil
}

func LargestValueN2(numbers []int) (*int, error) {
	if len(numbers) < 1 {
		return nil, ErrMinimumValue
	}
	maxValue := numbers[0]
	count := 0
	for i := 0; i < len(numbers); i++ {
		updated := false
		for j := i + 1; j < len(numbers); j++ {
			if maxValue < numbers[j] {
				maxValue = numbers[j]
				updated = true
			}
			count++
		}
		if !updated {
			break
		}
	}
	println("count: ", count)
	return intPtr(maxValue), nil
}

func intPtr(n int) *int {

	return &n
}

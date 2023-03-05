package applications

import (
	"errors"

	"github.com/stevedesilva/dsa-golang.git/commonsense/sort"
)

func FindGreatestProduct(array []int, numberRange int) (int, error) {
	if len(array) < 1 {
		return 0, errors.New("invalid array length")
	}
	if numberRange < 1 || len(array) < numberRange {
		return 0, errors.New("invalid number range")
	}

	q := sort.NewQuickSort(array)
	return findGreatestProduct(q.Quicksort(), numberRange), nil
}

func findGreatestProduct(sorted []int, numberRange int) int {
	count, sum := 0, 0
	for count < numberRange {
		lastElement := len(sorted) - 1
		sum += sorted[lastElement-count]
		count++
	}
	return sum
}

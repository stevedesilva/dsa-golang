package array

import (
	"errors"
	"fmt"
)

var errNoInputSuppliedForInventory = errors.New("no input arguments provided")

func markInventory(items []string) ([]string, error) {
	if len(items) <= 0 {
		return nil, errNoInputSuppliedForInventory
	}
	res := make([]string, 0, len(items))
	end := 5
	for _, i := range items {
		for j := 1; j <= end; j++ {
			res = append(res, fmt.Sprintf("%s Size: %d", i, j))
		}
	}
	return res, nil
}

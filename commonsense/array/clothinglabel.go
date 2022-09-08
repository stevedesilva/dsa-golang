package array

import "fmt"

func markInventory(items []string) ([]string, error) {
	res := make([]string, 0, len(items))
	end := 5
	for _, i := range items {
		for j := 1; j <= end; j++ {
			res = append(res, fmt.Sprintf("%s Size: %d", i, j))
		}
	}
	return res, nil
}

package optimise

import (
	"errors"
	"math"
)

func MaxProduct(nums []int) (int, error) {
	if len(nums) < 2 {
		return 0, errors.New("not enough numbers to form a product")
	}

	max1 := -math.MaxInt
	max2 := -math.MaxInt
	least1 := math.MaxInt
	least2 := math.MaxInt

	for _, num := range nums {
		if num > max1 {
			max2 = max1
			max1 = num
		} else if num > max2 {
			max2 = num
		}
		if num < least1 {
			least2 = least1
			least1 = num
		} else if num < least2 {
			least2 = num
		}

	}
	product1 := max1 * max2
	product2 := least1 * least2
	if product1 > product2 {
		return product1, nil
	} else {
		return product2, nil
	}
}

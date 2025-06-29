package optimise

import "errors"

func MaxNum(array []int) (int, error) {
	if len(array) == 0 {
		return 0, errors.New("Array cannot be empty")
	}
	if len(array) == 1 {
		return array[0], nil
	}
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	return max, nil

}

package optimise

import (
	"errors"
	"fmt"
)

var ErrTemperatureOutOfRange = errors.New("temperature out of range")

func SortBodyTemperatures(temp []float64) ([]float64, error) {
	// check all temperatures are in range
	temperatures := make(map[float64]int)
	for _, t := range temp {
		if t < 97.0 || t > 99.0 {
			return nil, ErrTemperatureOutOfRange
		}

		// if temperature already exists, increment count
		if count, exists := temperatures[t]; exists {
			temperatures[t] = count + 1
		} else {
			temperatures[t] = 1
		}
	}
	tempStart := 970
	tempEnd := 990
	// loop from tempStart to tempEnd and append to result if exists in map
	result := make([]float64, 0, len(temp))
	for t := tempStart; t <= tempEnd; t++ {
		fmt.Printf("%d\n", t)
		current := float64(t) / 10.0
		if count, exists := temperatures[current]; exists {
			for i := 0; i < count; i++ {
				result = append(result, current)
			}
		}
	}
	return result, nil
}

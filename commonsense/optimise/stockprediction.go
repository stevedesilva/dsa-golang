package optimise

import (
	"errors"
)

func FindGreatestProfit(prices []float64) (float64, error) {
	if len(prices) < 2 {
		return 0, errors.New("not enough prices to calculate profit")
	}

	low := prices[0]
	high := prices[0]
	total := 0.0

	for i := 0; i < len(prices); i++ {
		curr := prices[i]
		if curr < low {
			// set highest
			low = curr
			high = curr
		} else if curr > high {
			// set lowest
			high = curr
		}
		currTotal := high - low
		if total < currTotal {
			total = currTotal
		}
	}

	return total, nil
}

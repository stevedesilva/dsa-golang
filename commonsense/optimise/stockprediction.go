package optimise

import "errors"

func FindGreatestProfit(prices []float64) (float64, error) {

	if len(prices) < 2 {
		return 0, errors.New("not enough prices to calculate profit")
	}

	low := 0.0
	total := 0.0

	high := prices[0]
	for i := 1; i < len(prices); i++ {
		// get highest
		if prices[i] > high {
			high = prices[i]
		}
		// get lowest
		if prices[i] < low {
			low = prices[i]
		}
		total = high - low
	}

	return total, nil
}

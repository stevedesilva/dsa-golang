package optimise

import "errors"

func IsIncreasingTriplet(prices []float64) (bool, error) {
	if len(prices) < 3 {
		return false, errors.New("Array must contain at least three elements")
	}
	// Initialize the first two elements of the triplet
	lowestPrice := prices[0]
	// max value of a float64
	middlePrice := 1.7976931348623157e+308 // Maximum value for float64

	for _, price := range prices {
		if price <= lowestPrice {
			lowestPrice = price
		} else if price <= middlePrice {
			middlePrice = price
		} else {
			return true, nil
		}

	}

	return false, nil
}

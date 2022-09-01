package array

import "math"

func AverageCelsuisReader(fahrenheitReadings []int) int {
	celsuisNumbers := make([]int, 0, len(fahrenheitReadings))
	// convert fahrenheit to celsuis
	for _, f := range fahrenheitReadings {
		fahrenheitToCelsuis := math.Round((float64(f) - 32.0) * 0.5556)
		celsuisNumbers = append(celsuisNumbers, int(fahrenheitToCelsuis))
	}

	sum := 0
	for _, c := range celsuisNumbers {
		sum += c
	}
	avg := sum / len(celsuisNumbers)
	return avg

}

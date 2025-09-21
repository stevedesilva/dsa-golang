package optimise

import "errors"

var ErrTemperatureOutOfRange = errors.New("temperature out of range")

/*
Takes a [] of float64 representing body temperatures in Fahrenheit
Sorts the temperatures into ascending order.
return error if any temperature in the array is not in range 97.0 to 99.0 inclusive
*/
func SortBodyTemperatures(temp []float64) ([]float64, error) {
	// check all temperatures are in range
	for _, t := range temp {
		if t < 97.0 || t > 99.0 {
			return nil, ErrTemperatureOutOfRange
		}
	}
	return nil, nil
}

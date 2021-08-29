package section7

import (
	"strconv"
)

func FizzBuzz(num int) []string {
	result := make([]string, 0, num)
	for i := 1; i <= num; i++ {
		if i%5 == 0 && i%3 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

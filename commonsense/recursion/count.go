package recursion

import "fmt"

func countTo(number int) []int {
	res := countDown(number, make([]int, 0))
	return res
}

func countDown(number int, acc []int) []int {
	if number == 0 {
		// old slice
		return acc
	}
	fmt.Printf("num %d \t acc %v \n", number, acc)
	// new slice created when append
	acc = append(acc, number)
	return countDown(number-1, acc)
}

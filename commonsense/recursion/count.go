package recursion

import "fmt"

func CountTo(number int) []int {
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
	newAcc := append(acc, number)
	return countDown(number-1, newAcc)
}

func CountEven(low, high int) []int {
	res := make([]int, 0)
	return countEvenRec(low, high, res)
}

func countEvenRec(low, high int, acc []int) []int {
	if low > high {
		return acc
	}
	acc = append(acc, low)
	return countEvenRec(low+2, high, acc)
}

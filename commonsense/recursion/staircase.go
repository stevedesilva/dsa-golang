package recursion

import "fmt"

func CountNumberOfSteps(number int) int {
	fmt.Println(">", number)
	res := countNumberOfSteps(number)
	fmt.Println("<", res)
	return res
}

//func countNumberOfSteps(number int) int {
//
//	if number <= 0 {
//		return 0
//	}
//	if number == 1 {
//		return 1
//	}
//	if number == 2 {
//		return 2
//	}
//	if number == 3 {
//		return 4
//	}
//
//	fmt.Println(":", number)
//	return countNumberOfSteps(number-1) + countNumberOfSteps(number-2) + countNumberOfSteps(number-3)
//}

func countNumberOfSteps(number int) int {

	if number <= 0 {
		return 0
	}
	if number == 1 {
		return 1
	}
	if number == 2 {
		return 2
	}
	if number == 3 {
		return 4
	}

	fmt.Println(":", number)
	return countNumberOfSteps(number - 1)
}

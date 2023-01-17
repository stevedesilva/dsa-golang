package dynamic_programming

import (
	"fmt"
)

func Fibonacci(number int) int {
	cache := make(map[int]int, 0)
	return fibonacci(number, cache)
}

func fibonacci(number int, cache map[int]int) int {
	if number <= 1 {
		return number
	}
	if _, found := cache[number]; !found {
		cache[number] = fibonacci(number-2, cache) + fibonacci(number-1, cache)
		fmt.Println(cache[number])
		return cache[number]
	}
	return cache[number]
}

func FibonacciNoRecursive(number int) int {
	if number <= 1 {
		return number
	}
	sequence := []int{0, 1}
	for i := 2; i <= number; i++ {
		res := sequence[i-2] + sequence[i-1]
		sequence = append(sequence, res)
	}
	return sequence[len(sequence)-1]
}

// 0 1 1 2 3 5 8 13 21
func FibonacciNoRecursiveAlt(number int) int {
	a := 0
	b := 1
	if number <= 1 {
		return number
	}
	temp := 0
	for i := 2; i <= number; i++ {
		temp = a + b
		a = b
		b = temp
	}

	return temp
}

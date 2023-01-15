package dynamic_programming

import "fmt"

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
	return 0
}

package dynamic_programming

import "fmt"

func Golomb(number int) int {
	cache := make(map[int]int)
	return golomb(number, cache)
}

func golomb(number int, cache map[int]int) int {
	if number == 1 {
		return 1
	}
	if val, ok := cache[number]; ok {
		fmt.Printf("cached number%d= %d \n", number, val)
		return val
	}
	cache[number] = 1 + golomb(number-golomb(golomb(number-1, cache), cache), cache)
	return cache[number]
}

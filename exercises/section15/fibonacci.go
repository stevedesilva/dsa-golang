package section15

import "fmt"

func GenerateFibSeriesIterative(num int) int {
	res := []int{0, 1}
	for i := 2; i <= num; i++ {
		value := res[(len(res)-2)] + res[(len(res)-1)]
		res = append(res, value)
	}
	return res[(len(res) - 1)]
}

func GenerateFibSeriesIterative2(num int) int {
	res := []int{0, 1}
	for i := 2; i <= num; i++ {
		value := res[i-2] + res[i-1]
		res = append(res, value)
	}
	return res[num]
}

func GenerateFibSeriesRecursive(num int) int {
	res := fibSeriesRecursive(num, 2, []int{0, 1})
	return res[(len(res) - 1)]
}

func fibSeriesRecursive(num, count int, res []int) []int {
	if count > num {
		return res
	}

	value := res[(len(res)-2)] + res[(len(res)-1)]
	res = append(res, value)

	count++
	return fibSeriesRecursive(num, count, res)
}

func GenerateFibSeriesRecursive2(num int) int {
	res := fibSeriesRecursive2(num, 2, []int{0, 1})
	return res[num]
}

func fibSeriesRecursive2(num, count int, res []int) []int {
	if count > num {
		return res
	}

	value := res[count-2] + res[count-1]
	res = append(res, value)

	count++
	return fibSeriesRecursive(num, count, res)
}

// 0 1 1 2 3 5 8 13 21 34
func FibRecursiveNaive(num int) int {
	if num < 2 {
		return num
	}
	return FibRecursiveNaive(num-1) + FibRecursiveNaive(num-2)
}

func slowFib(num int, cache map[int]int) int {
	if val, ok := cache[num]; ok {
		fmt.Printf("cache hit: %d\n", val)
		return val
	} else {
		if num < 2 {
			cache[num] = num
			fmt.Printf("cache add base: %d\n", num)
			return num
		}
		//numLessTwo := num - 2
		//second := slowFib(numLessTwo, cache)
		//cache[numLessTwo] = second
		//fmt.Printf("cache add second: key(%d) val(%d)\n", numLessTwo, second)
		//
		//numLessOne := num - 1
		//first := slowFib(numLessOne, cache)
		//cache[numLessOne] = first
		//fmt.Printf("cache add first: key(%d) val(%d)\n", numLessOne, first)

		val := slowFib(num-2, cache) + slowFib(num-1, cache)
		cache[num] = val
		fmt.Printf("cache add: key(%d) val(%d)\n", num, val)
		return cache[num]
	}
}

func FibRecursiveMemoizedV1(num int) int {
	cache := make(map[int]int)
	result := slowFib(num, cache)
	return result
}

func FibRecursiveMemoizedV2(num int) int {
	fastFib := newMemoizedFib()
	result := fastFib(num)
	return result
}

func newMemoizedFib() func(int) int {
	cache := make(map[int]int)
	var fn func(int) int
	fn = func(n int) int {
		if n == 1 || n == 2 {
			return 1
		}
		if _, ok := cache[n]; !ok {
			cache[n] = fn(n-1) + fn(n-2)
		}
		return cache[n]
	}
	return fn
}

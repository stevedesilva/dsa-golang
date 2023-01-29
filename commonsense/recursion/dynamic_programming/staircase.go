package dynamic_programming

func CountNumberOfPathsWithMemoization(steps int) int {
	cache := make(map[int]int)
	return countNumberOfPathsWithMemoization(steps, cache)
}

func countNumberOfPathsWithMemoization(steps int, cache map[int]int) int {
	if steps <= 0 {
		return 0
	}
	if steps == 1 {
		return 1
	}
	if steps == 2 {
		return 2
	}
	if steps == 3 {
		return 4
	}
	// 8 = 7r(44) + 6r(24) + 5r(13)
	// 7 = 6r + 5r + 4r  = 24 + 13 + 7 = 44
	// 6 = 5r + 4r + 3(4) = 13 + 7 + 4 = 24
	// 5 = 4r + 3(4) + 2(2) = 7 + 4 + 2 = 13
	// 4 = 3(4) + 2(2) + 1(1) = 7
	// 3 = 4
	// 2 = 2
	// 1 = 1
	if val, ok := cache[steps]; !ok {
		val = countNumberOfPathsWithMemoization(steps-1, cache) + countNumberOfPathsWithMemoization(steps-2, cache) + countNumberOfPathsWithMemoization(steps-3, cache)
		cache[steps] = val
	}
	return cache[steps]
}

func CountNumberOfPathsAltWithMemoization(steps int) int {

	// 8 = 7r(44) + 6r(24) + 5r(13)
	// 7 = 6r + 5r + 4r  = 24 + 13 + 7 = 44
	// 6 = 5r + 4r + 3(4) = 13 + 7 + 4 = 24
	// 5 = 4r + 3(4) + 2(2) = 7 + 4 + 2 = 13
	// 4 = 3(4) + 2(2) + 1(1) = 7
	// 3 = 4
	// 2 = 2
	// 1 = 1
	cache := make(map[int]int)
	return countNumberOfPathsAltWithMemoization(steps, cache)
}

func countNumberOfPathsAltWithMemoization(steps int, cache map[int]int) int {
	// hardcode base values
	if steps < 0 {
		return 0
	}
	if steps == 0 || steps == 1 {
		return 1
	}
	if val, ok := cache[steps]; !ok {
		val = countNumberOfPathsAltWithMemoization(steps-1, cache) + countNumberOfPathsAltWithMemoization(steps-2, cache) + countNumberOfPathsAltWithMemoization(steps-3, cache)
		cache[steps] = val
	}
	return cache[steps]
}

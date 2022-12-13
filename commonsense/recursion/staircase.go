package recursion

func CountNumberOfPaths(steps int) int {
	// hardcode base values
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

	return CountNumberOfPaths(steps-1) + CountNumberOfPaths(steps-2) + CountNumberOfPaths(steps-3)
}

func CountNumberOfPathsAlt(steps int) int {
	// hardcode base values
	if steps < 0 {
		return 0
	}
	if steps == 0 || steps == 1 {
		return 1
	}
	// 8 = 7r(44) + 6r(24) + 5r(13)
	// 7 = 6r + 5r + 4r  = 24 + 13 + 7 = 44
	// 6 = 5r + 4r + 3(4) = 13 + 7 + 4 = 24
	// 5 = 4r + 3(4) + 2(2) = 7 + 4 + 2 = 13
	// 4 = 3(4) + 2(2) + 1(1) = 7
	// 3 = 4
	// 2 = 2
	// 1 = 1
	return CountNumberOfPathsAlt(steps-1) + CountNumberOfPathsAlt(steps-2) + CountNumberOfPathsAlt(steps-3)
}

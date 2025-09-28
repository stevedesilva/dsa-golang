package optimise

func LongestSequence(nums []int) []int {
	// return longest consecutive sequence in nums
	numMap := make(map[int]bool, len(nums))
	// add numbers to map
	for _, v := range nums {
		numMap[v] = true
	}
	// loop over numbers map - check if num-1 exists in map
	result := []int{}
	for _, num := range nums {
		println(num)
		currRes := []int{}
		if !numMap[num-1] {
			// add current to map
			currRes = append(currRes, num)
			// increment
			curr := num + 1
			for numMap[curr] {
				// add new number to list
				currRes = append(currRes, curr)
				curr++
			}
			if len(currRes) > len(result) {
				result = currRes
			}
			println(currRes)
		}
	}
	return result
}

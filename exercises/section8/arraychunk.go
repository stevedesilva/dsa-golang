package section8

func ChunkIntoSlices(nums []int, chunkSize int) [][]int {
	res := make([][]int, 0)
	numLen := len(nums)

	for startIdx := 0; startIdx < numLen; startIdx = startIdx + chunkSize {
		// end Index
		endIdx := startIdx + chunkSize
		if endIdx > numLen {
			endIdx = numLen
		}
		// add to array
		res = append(res, nums[startIdx:endIdx])

	}
	return res
}

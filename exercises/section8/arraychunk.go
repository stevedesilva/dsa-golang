package section8

func ChunkIntoSlices(nums []int, chunkSize int) [][]int {
	res := make([][]int, 0)
	numLen := len(nums)

	for startIdx := 0; startIdx < len(nums); startIdx = startIdx + chunkSize {
		// get next chunk
		if startIdx >= numLen {
			break
		}
		// reset length
		endIdx := startIdx + chunkSize
		if endIdx > numLen {
			endIdx = numLen
		}
		// add to array
		res = append(res, nums[startIdx:endIdx])

	}
	return res
}

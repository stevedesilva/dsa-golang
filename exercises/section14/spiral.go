package section14

func CreateSpiral(size int) [][]int {

	spiral := make([][]int, size)
	for i := range spiral {
		spiral[i] = make([]int, size)
	}

	return spiral

}

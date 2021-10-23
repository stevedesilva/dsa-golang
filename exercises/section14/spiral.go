package section14

func CreateSpiral(size int) [][]int {
	// initialise array
	spiral := make([][]int, size)
	for i := range spiral {
		spiral[i] = make([]int, size)
	}

	startRow := 0
	startColumn := 0
	endRow := size - 1
	endColumn := size - 1

	count := 1
	for startRow <= endRow && startColumn <= endColumn {
		// (row-static)(column->) TOP ROW: top left -> top right
		for i := startColumn; i <= endColumn; i++ {
			spiral[startRow][i] = count
			count++
		}
		startRow++
		// (row-v)(column-static) RIGHT COLUMN: top right -> bottom right
		for i := startRow; i <= endRow; i++ {
			spiral[i][endColumn] = count
			count++
		}
		endColumn--
		// (row-static)(column<-) BOTTOM ROW: bottom right -> bottom left
		for i := endColumn; i >= startColumn; i-- {
			spiral[endRow][i] = count
			count++
		}
		endRow--
		// (row-static)(column-^) LEFT COLUMN: bottom left -> top left
		for i := endRow; i >= startRow; i-- {
			spiral[i][startColumn] = count
			count++
		}
		startColumn++
	}

	return spiral

}

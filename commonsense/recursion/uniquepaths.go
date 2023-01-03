package recursion

/*
	C  C  C  C  C

R 70 35 15  5  1
R 35 20 10  4  1
R 15 10  6  3  1
R  5  4  3  2  1
R  1  1  1  1  1
*/
func findAllUniquePaths(row, column int) int {
	if row == 1 || column == 1 {
		return 1
	}
	return findAllUniquePaths(row-1, column) + findAllUniquePaths(row, column-1)
}

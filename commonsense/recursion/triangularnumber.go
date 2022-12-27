package recursion

func CalculateTriangularNumber(positionToCalculate int) int {
	return calculateTriangularNumberRec(positionToCalculate, 0, 0)
}

func calculateTriangularNumberRec(positionToCalculate, currentIndex, currentValue int) int {
	if currentIndex > positionToCalculate {
		return currentValue
	}
	return calculateTriangularNumberRec(positionToCalculate, currentIndex+1, currentValue+currentIndex)
}

func CalculateTriangularNumberAlt(n int) int {
	if n == 1 {
		return 1
	}
	return n + CalculateTriangularNumberAlt(n-1)
}

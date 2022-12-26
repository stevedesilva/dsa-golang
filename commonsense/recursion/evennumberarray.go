package recursion

func EvenNumbersFromArray(array []int) []int {
	res := make([]int, 0)
	return evenNumbersFromArray(array, res)

}

func evenNumbersFromArray(array, acc []int) []int {
	if len(array) == 0 {
		return acc
	}
	if even := array[0]%2 == 0; even {
		acc = append(acc, array[0])
	}
	return evenNumbersFromArray(array[1:], acc)
}

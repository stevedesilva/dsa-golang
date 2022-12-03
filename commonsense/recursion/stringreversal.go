package recursion

func StringReversal(word string) string {
	res := stringReversalRec([]rune(word))
	return string(res)
}

func stringReversalRec(array []rune) string {
	if len(array) == 1 {
		return string(array[0])
	}
	return stringReversalRec(array[1:]) + string(array[0])
}

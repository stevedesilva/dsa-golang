package section3

//func ReverseInlineOptimal(s string) string {
//	r := []rune(s)
//	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
//		r[i], r[j] = r[j], r[i]
//	}
//	return string(r)
//}

func ReverseInlineOptimal(input string) string {
	inRune := []rune(input)
	for i, j := 0, len(inRune)-1; i < j; i, j = i+1, j-1 {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	}
	return string(inRune)
}

func ReverseWithResultArray(input string) string {
	inputAsRune := []rune(input)
	// Gotcha : set length with a capacity when using append - otherwise it will add new elements to the end of the slice
	result := make([]rune, 0, len(inputAsRune))

	for i := len(inputAsRune) - 1; i >= 0; i-- {
		result = append(result, inputAsRune[i])
	}
	return string(result)
}

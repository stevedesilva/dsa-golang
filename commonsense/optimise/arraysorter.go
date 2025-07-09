package optimise

func SortCharactersInAnyOrder(characters []rune) []rune {
	charMap := make(map[rune]int)
	sorted := make([]rune, 0, len(characters))
	// add chars into a map of rune int count
	for _, v := range characters {
		charMap[v]++
	}

	// create list of sorted characters
	for key, count := range charMap {
		for i := 0; i < count; i++ {
			sorted = append(sorted, key)
		}
	}

	return sorted
}

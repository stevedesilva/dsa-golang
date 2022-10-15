package hashmap

func findFirstMissingCharacter(sentence string) string {
	alphabet := make(map[rune]bool)
	for _, c := range sentence {
		alphabet[c] = true
	}

	var missing string
	for i := 'a'; i <= 'z'; i++ {
		if !alphabet[i] {
			missing = string(i)
			break
		}
		alphabet[i] = false
	}

	return missing
}

func findMissingCharacters(sentence string) []string {
	alphabet := make(map[rune]bool)
	for _, c := range sentence {
		alphabet[c] = true
	}

	var missing = make([]string, 0)
	for i := 'a'; i <= 'z'; i++ {
		if !alphabet[i] {
			missing = append(missing, string(i))
			break
		}
		alphabet[i] = false
	}

	return missing
}

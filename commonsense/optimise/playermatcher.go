package optimise

type PlayerMatcher struct {
	firstName string
	lastName  string
	team      string
}

type PlayerResult struct {
	firstName string
	lastName  string
}

func IsPlayerMatch(playersA, playersB []PlayerMatcher) []PlayerResult {
	// validation
	if playersA == nil || playersB == nil {
		return []PlayerResult{}
	}
	// add to set playerA
	playersASet := make(map[PlayerMatcher]bool)
	for _, playerA := range playersA {
		playersASet[playerA] = true
	}
	// add to set playerB
	playersBSet := make(map[PlayerMatcher]bool)
	for _, playerB := range playersB {
		playersBSet[playerB] = true
	}
	// compare set
	results := make([]PlayerResult, 0, len(playersASet))
	for key := range playersASet {
		if playersBSet[key] {
			results = append(results, PlayerResult{key.firstName, key.lastName})
		}
	}
	//alternative
	//for key, _ := range playersASet {
	//	if _, found := playersBSet[key]; found {
	//		results = append(results, PlayerResult{key.firstName, key.lastName})
	//	}
	//}
	// build return list
	return results
}

package optimise

import "strings"

const ME = "ME"
const THEM = "THEM"

func flipCoin(coinNumber int, currentPlayer string) string {

	if (coinNumber-1)%3 == 0 {
		if strings.EqualFold(currentPlayer, ME) {
			return THEM
		} else {
			return ME
		}

	}
	if strings.EqualFold(currentPlayer, ME) {
		return ME
	} else {
		return THEM
	}
}

package optimise

func flipCoin(coinNumber int, currentPlayer string) string {
	if (coinNumber-1)%3 == 0 {
		return "THEM"
	}
	return "ME"
}

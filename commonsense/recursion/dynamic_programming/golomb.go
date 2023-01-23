package dynamic_programming

func Golomb(number int) int {
	if number == 1 {
		return 1
	}
	return 1 + Golomb(number-Golomb(Golomb(number-1)))
}

package array

import "strconv"

func IsCreditCardValidAsNumber(creditCardNumber int) bool {
	return IsCreditCardValidAsString(strconv.Itoa(creditCardNumber))
}

func IsCreditCardValidAsString(creditCardNumber string) bool {
	// double every number from right to left
	workingNum := []rune(creditCardNumber)

	flip := false
	for i := len(workingNum) - 1; i >= 0; i-- {
		curr := asciiToInt(workingNum[i])

		if flip {
			// if number is > 9 then split and add numbers together
			curr *= 2
			if curr > 9 {
				digits := []rune(strconv.Itoa(int(curr)))
				curr = asciiToInt(digits[0]) + asciiToInt(digits[1])
			}
			workingNum[i] = curr
		} else {
			workingNum[i] = curr
		}
		flip = !flip
	}

	// sum number
	var sum rune
	for _, v := range workingNum {
		sum += v
	}
	// mod by 10 return true if number is 0
	return (sum % 10) == 0
}

func asciiToInt(asciiValue rune) rune {
	return asciiValue - '0'
}

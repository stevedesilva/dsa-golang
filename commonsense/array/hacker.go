package array

import "math"

type PasswordType struct {
	accumulator []string
}

// https://www.geeksforgeeks.org/print-all-combinations-of-given-length/
func (p *PasswordType) PasswordCracker(set []rune, desiredLen int) []string {
	if len(set) <= 0 {
		return make([]string, 0)
	}
	// x to the power y
	pow := math.Pow(float64(len(set)), float64(desiredLen))
	p.accumulator = make([]string, 0, int(pow))
	p.passwordCrackerRec(set, "", desiredLen)
	return p.accumulator
}

func (p *PasswordType) passwordCrackerRec(set []rune, prefix string, desiredLen int) {
	// base case - we have got our desired length
	if desiredLen == 0 {
		p.accumulator = append(p.accumulator, prefix)
		return
	}
	// for every char desired set find all possible combinations
	for i := 0; i < len(set); i++ {
		nextChar := prefix + string(set[i])
		p.passwordCrackerRec(set, nextChar, desiredLen-1)
	}

}

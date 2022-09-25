package array

// https://www.geeksforgeeks.org/print-all-combinations-of-given-length/
func PasswordCracker(set []rune, desiredLen int) []string {
	i := len(set) * desiredLen
	accumulator := make([]string, 0, i)
	passwordCrackerRec(set, len(set), "", desiredLen, accumulator)
	// convert acc
	return accumulator
}

func passwordCrackerRec(set []rune, setLen int, prefix string, desiredLen int, accumulator []string) {

}

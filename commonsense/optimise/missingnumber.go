package optimise

import "errors"

func FindAlternative(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty list")
	}
	numberSet := make(map[int]bool)
	for _, n := range numbers {
		numberSet[n] = true
	}
	for i := 0; i < len(numbers); i++ {
		if !numberSet[i] {
			return i, nil
		}
	}

	// if number not found, return error
	return 0, errors.New("no missing number found")
}

func Find(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("empty list")
	}

	// Calculate the expected sum of numbers from 0 to n
	n := len(numbers)
	expectedSum := n * (n + 1) / 2

	// Calculate the actual sum of the given numbers
	actualSum := 0
	for _, num := range numbers {
		actualSum += num
	}

	// The missing number is the difference between the expected and actual sums
	missingNumber := expectedSum - actualSum

	// Validate if the missing number is within the valid range
	if missingNumber < 0 || missingNumber > n {
		return 0, errors.New("no missing number found")
	}

	return missingNumber, nil
}

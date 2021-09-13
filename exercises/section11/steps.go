package section11

import (
	"fmt"
	"strings"
)

func ExecuteStep(steps int) []string {
	result := make([]string, 0, steps)
	for i := 1; i <= steps; i++ {
		s := strings.Repeat("#", i)
		space := strings.Repeat(" ", steps-i)
		step := fmt.Sprintf("%s%s", s, space)
		result = append(result, step)
	}
	PrintRows(result)
	return result
}

func ExecuteStepNaive(steps int) []string {
	result := make([]string, 0, steps)
	for row := 0; row < steps; row++ {
		current := make([]string, 0, steps)
		for column := 0; column < steps; column++ {
			if column <= row {
				current = append(current, "#")
			} else {
				current = append(current, " ")
			}
		}
		cRow := strings.Join(current, "")
		result = append(result, cRow)
	}
	PrintRows(result)
	return result
}

func PrintRows(rows []string) {
	for _, cRow := range rows {
		fmt.Println(cRow)
	}
}

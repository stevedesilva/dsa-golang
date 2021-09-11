package section11

import (
	"fmt"
	"strings"
)

func executeStep(steps int) []string {
	result := make([]string, 0, steps)
	for i := 1; i <= steps; i++ {
		s := strings.Repeat("#", i)
		space := strings.Repeat(" ", steps-i)
		step := fmt.Sprintf("%s%s\n", s, space)
		result = append(result, step)
	}
	return result
}

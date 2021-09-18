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

func ExecuteRecursive(step int) []string {
	res := make([]string, step)
	col := make([]string, step)
	rowNum := 0
	executeRecursive(step, rowNum, col, res, 0)
	return res
}

func executeRecursive(steps int, rowNum int, col []string, result []string, colNum int) {
	if steps == rowNum {
		return
	}

	if steps == colNum {
		result[rowNum] = strings.Join(col, "")
		newCol := make([]string, steps)
		executeRecursive(steps, rowNum+1, newCol, result, 0)
		return
	}

	if colNum <= rowNum {
		col[colNum] = "#"
	} else {
		col[colNum] = " "
	}
	executeRecursive(steps, rowNum, col, result, colNum+1)
}

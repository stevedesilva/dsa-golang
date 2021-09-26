package section12

import (
	"fmt"
	"strings"
)

type Pyramid struct {
	level int
}

func (p *Pyramid) BuildPyramid() []string {
	result := make([]string, 0, p.level)
	halfLevel := p.level - 1
	colMax := p.level + p.level - 1
	for row := 0; row < p.level; row++ {
		rowArray := make([]string, 0, p.level)
		left := halfLevel - row
		right := halfLevel + row
		for col := 0; col < colMax; col++ {
			if col < left || col > right {
				rowArray = append(rowArray, " ")
			} else {
				rowArray = append(rowArray, "#")
			}
		}
		rowAsStr := strings.Join(rowArray, "")
		result = append(result, rowAsStr)
	}
	PrintPyramid(result)
	return result
}

func PrintPyramid(rows []string) {
	for _, row := range rows {
		fmt.Println(row)
	}
}

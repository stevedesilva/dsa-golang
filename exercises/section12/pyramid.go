package section12

import (
	"fmt"
	"strings"
)

type Pyramid struct {
	level int
}

func (p *Pyramid) BuildPyramidIterative() []string {
	pyramid := make([]string, 0, p.level)
	colHalf := p.level - 1
	colMax := p.level + p.level - 1

	for currentLevel := 0; currentLevel < p.level; currentLevel++ {
		levelSlice := make([]string, 0, p.level)
		for col := 0; col < colMax; col++ {
			if col < colHalf-currentLevel || col > colHalf+currentLevel {
				levelSlice = append(levelSlice, " ")
			} else {
				levelSlice = append(levelSlice, "#")
			}
		}
		rowAsStr := strings.Join(levelSlice, "")
		pyramid = append(pyramid, rowAsStr)
	}

	PrintPyramid(pyramid)
	return pyramid
}

func PrintPyramid(rows []string) {
	for _, row := range rows {
		fmt.Println(row)
	}
}

func (p *Pyramid) BuildPyramidRecursive(maxLevel int) []string {
	pyramid := make([]string, maxLevel)
	newLevelSlice := make([]string, maxLevel)
	rowNum := 0
	colNum := 0
	colHalf := maxLevel - 1
	colMax := maxLevel + maxLevel - 1
	executeRecursive(maxLevel, rowNum, newLevelSlice, pyramid, colNum, colHalf, colMax)
	return pyramid
}

func executeRecursive(maxLevel int,
	currentLevel int,
	levelSlice []string,
	pyramid []string,
	colNum int,
	colHalf int,
	colMax int) {

	if maxLevel == currentLevel {
		return
	}

	if colNum == colMax {
		pyramid[currentLevel] = strings.Join(levelSlice, "")
		newLevelSlice := make([]string, maxLevel)
		executeRecursive(maxLevel, currentLevel+1, newLevelSlice, pyramid, 0, colHalf, colMax)
		return
	}

	if colNum < colHalf-currentLevel || colNum > colHalf+currentLevel {
		levelSlice = append(levelSlice, " ")
	} else {
		levelSlice = append(levelSlice, "#")
	}

	executeRecursive(maxLevel, currentLevel, levelSlice, pyramid, colNum+1, colHalf, colMax)
}

package section4

import (
	"fmt"
	"strconv"
)

func Reverse(input int) int {
	inAsString := strconv.Itoa(input)
	inputAsRunes := []rune(inAsString)

	i, j := 0, len(inputAsRunes)-1
	if inputAsRunes[0] == 45 {
		i = 1
	}
	for ; i < j; i, j = i+1, j-1 {
		fmt.Printf("Loop %d %d %v %v\n", i, j, inputAsRunes[i], inputAsRunes[j])
		inputAsRunes[i], inputAsRunes[j] = inputAsRunes[j], inputAsRunes[i]
	}

	res, _ := strconv.Atoi(string(inputAsRunes))
	//if err != nil {
	//
	//}
	return res
}

//func Reverse(input int) int {
//	itoa := strconv.Itoa(input)
//	//fmt.Printf("%T %[1]v\n", itoa)
//	toInt := []rune(itoa)
//	//fmt.Printf("%T %[1]v\n", toInt)
//
//	for i, j := 0, len(toInt)-1; i > j; i, j = i+1, j-1 {
//		fmt.Printf("Loop %d %d %v %v\n", i, j, toInt[i], toInt[j])
//		//inputAsRunes[i][j] = inputAsRunes[j][i]
//	}
//
//	res, _ := strconv.Atoi(string(toInt))
//	//if err != nil {
//	//
//	//}
//	return res
//
//}

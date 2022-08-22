package array

import "fmt"

func HasIntersection(arrayA []int, arrayB []int) []int {
	step, count := 0, 0
	intersection := make([]int, 0)
	for i := 0; i < len(arrayA); i++ {
		for j := 0; j < len(arrayB); j++ {
			count++
			if arrayA[i] == arrayB[j] {
				step++
				intersection = append(intersection, arrayA[i])
				break
			}
		}
	}
	fmt.Println("Comparison number: ", step)
	fmt.Println("Total Count: ", count)
	fmt.Println("Result array: ", intersection)
	return intersection
}

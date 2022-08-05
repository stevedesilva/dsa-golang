package sort

import "fmt"

type BubbleSort struct {
	data []int
}

func (b *BubbleSort) Sort() {
	swapOccurred := false
	end := len(b.data)
	for i := 0; i < end; {
		fmt.Println("i : ", i)
		for j := i + 1; j < end; j++ {
			if b.data[j-1] > b.data[j] {
				b.data[j], b.data[j-1] = b.data[j-1], b.data[j]
				swapOccurred = true
			}
		}
		end--
		fmt.Println("end: ", end)
		fmt.Println("continue: ", i < end)
		if !swapOccurred {
			return
		}
	}
}

func (b *BubbleSort) SortAlt() {
	swapOccurred := false
	end := len(b.data)
	for i := 0; i < end; {
		fmt.Println("i : ", i)
		for j := i + 1; j < end; j++ {
			if b.data[j-1] > b.data[j] {
				b.data[j], b.data[j-1] = b.data[j-1], b.data[j]
				swapOccurred = true
			}
		}
		end--
		fmt.Println("end: ", end)
		fmt.Println("continue: ", i < end)
		if !swapOccurred {
			return
		}
	}
}

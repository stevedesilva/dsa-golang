package sort

type BubbleSort struct {
	data []int
}

func (b *BubbleSort) Sort() {
	swapOccurred := false
	end := len(b.data)
	for i := 0; i < end; {
		for j := i + 1; j < end; j++ {
			if b.data[j-1] > b.data[j] {
				b.data[j], b.data[j-1] = b.data[j-1], b.data[j]
				swapOccurred = true
			}
		}
		end--
		if !swapOccurred {
			return
		}
	}
}

func (b *BubbleSort) SortOfficial() {
	end := len(b.data)
	for swapped := true; swapped; {
		swapped = false
		for j := 1; j < end; j++ {
			if b.data[j] < b.data[j-1] {
				b.data[j], b.data[j-1] = b.data[j-1], b.data[j]
				swapped = true
			}
		}
		end--
	}
}

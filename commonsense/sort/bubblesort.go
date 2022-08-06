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
	swapped := true
	end := len(b.data)
	for swapped {
		swapped = false
		for j := 0; j < end; j++ {
			if b.data[j] > b.data[j+1] {
				b.data[j], b.data[j+1] = b.data[j+1], b.data[j]
				swapped = true
			}
		}
		end--
	}
}

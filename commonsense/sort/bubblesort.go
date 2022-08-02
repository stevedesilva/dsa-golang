package sort

type BubbleSort struct {
	data []int
}

func (b *BubbleSort) Sort() {
	for i := 0; i < len(b.data); i++ {
		for j := i + 1; j < len(b.data); j++ {
			if b.data[i] > b.data[j] {
				b.data[j], b.data[i] = b.data[i], b.data[j]
			}
		}
	}
}

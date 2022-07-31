package sort

type BubbleSort struct {
	data []int
}

func (b *BubbleSort) Sort() {
	for i, d := range b.data {
		_, _ = i, d
	}
}

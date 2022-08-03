package sort

type BubbleSort struct {
	data []int
}

//func (b *BubbleSort) Sort() {
//	//swapped := f
//	for i := 0; i < len(b.data); i++ {
//		for j := i + 1; j < len(b.data); j++ {
//			if b.data[i] > b.data[j] {
//				b.data[j], b.data[i] = b.data[i], b.data[j]
//			}
//		}
//	}
//}

func (b *BubbleSort) Sort() {
	notSwapped := false
	start := 0
	end := len(b.data)
	for i := start; i < end; i++ {
		if !notSwapped {
			return
		}
		for j := start + 1; j < end; j++ {
			if b.data[j-1] > b.data[j] {
				b.data[j], b.data[j-1] = b.data[j-1], b.data[j]
				notSwapped = true
			}
		}
		end--
	}
}

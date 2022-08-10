package sort

type SelectionSort struct {
	data []int
}

func (s *SelectionSort) Sort() {
	for i := 0; i < len(s.data); i++ {
		lowestIdx := i
		for j := i + 1; j < len(s.data); j++ {
			if s.data[j] < s.data[lowestIdx] {
				lowestIdx = j
			}
		}
		if s.data[lowestIdx] < s.data[i] {
			s.data[lowestIdx], s.data[i] = s.data[i], s.data[lowestIdx]
		}
	}
}

func (s *SelectionSort) SortOfficial() {
	for i := 0; i < len(s.data); i++ {
		lowestIdx := i
		for j := i + 1; j < len(s.data); j++ {
			if s.data[j] < s.data[lowestIdx] {
				lowestIdx = j
			}
		}
		if lowestIdx != i {
			s.data[lowestIdx], s.data[i] = s.data[i], s.data[lowestIdx]
		}
	}
}

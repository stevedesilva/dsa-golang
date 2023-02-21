package sort

type SortedArray struct {
	array []int
}

func NewQuickSelect(array []int) *SortedArray {
	return &SortedArray{
		array: array,
	}
}

func (q *SortedArray) QuickSelect() []int {
	q.quickselect(0, len(q.array)-1)
	return q.array
}

func (q *SortedArray) quickselect(leftIndex, rightIndex int) {
	if rightIndex-leftIndex <= 0 {
		return
	}

	pivotIndex := q.partition(leftIndex, rightIndex)

	q.quickselect(leftIndex, pivotIndex-1)
	q.quickselect(pivotIndex+1, rightIndex)
}

func (q *SortedArray) partition(leftPointer, rightPointer int) int {
	pivotIndex := rightPointer
	pivotValue := q.array[pivotIndex]

	rightPointer--
	for {
		for leftPointer >= 0 && q.array[leftPointer] < pivotValue {
			leftPointer++
		}

		for rightPointer >= 0 && q.array[rightPointer] > pivotValue {
			rightPointer--
		}

		if leftPointer >= rightPointer {
			break
		} else {
			q.array[leftPointer], q.array[rightPointer] = q.array[rightPointer], q.array[leftPointer]
			leftPointer++
		}
	}
	// swap pivot
	q.array[leftPointer], q.array[pivotIndex] = q.array[pivotIndex], q.array[leftPointer]

	return leftPointer

}

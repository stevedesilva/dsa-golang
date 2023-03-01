package sort

//func Quicksort(array []int) []int {
//	return array
//}

type SortableArray struct {
	array []int
}

func NewQuickSort(array []int) *SortableArray {
	return &SortableArray{
		array: array,
	}
}

func (q *SortableArray) Quicksort() []int {
	q.quicksort(0, len(q.array)-1)
	return q.array
}

func (q *SortableArray) quicksort(leftIndex, rightIndex int) {
	if rightIndex-leftIndex <= 0 {
		return
	}

	pivotIndex := q.partition(leftIndex, rightIndex)

	q.quicksort(leftIndex, pivotIndex-1)
	q.quicksort(pivotIndex+1, rightIndex)
}

func (q *SortableArray) partition(leftPointer, rightPointer int) int {
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

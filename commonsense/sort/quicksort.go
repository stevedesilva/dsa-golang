package sort

func Quicksort(array []int) []int {
	return array
}

type SortableArray struct {
	array []int
}

func New() *SortableArray {
	return &SortableArray{}
}

func (q *SortableArray) partition(leftPointer, rightPointer int) int {
	pivotIndex := rightPointer
	pivotValue := q.array[pivotIndex]

	rightPointer--
	for {
		for q.array[leftPointer] < pivotValue {
			leftPointer++
		}

		for q.array[rightPointer] > pivotValue {
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

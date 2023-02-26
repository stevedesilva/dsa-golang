package sort

type SortedArray struct {
	array []int
}

func NewQuickSelect(array []int) *SortedArray {
	return &SortedArray{
		array: array,
	}
}

func (q *SortedArray) QuickSelect(indexToFind int) int {
	return q.quickselect(indexToFind, 0, len(q.array)-1)
}

func (q *SortedArray) quickselect(indexToFind, leftIndex, rightIndex int) int {
	if rightIndex-leftIndex <= 0 {
		return q.array[leftIndex]
	}

	partition := q.partition(leftIndex, rightIndex)
	if indexToFind < partition {
		return q.quickselect(indexToFind, leftIndex, partition-1)

	} else if indexToFind > partition {
		return q.quickselect(indexToFind, partition+1, rightIndex)
	} else {
		return q.array[indexToFind]
	}
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

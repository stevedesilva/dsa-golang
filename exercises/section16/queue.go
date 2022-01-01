package section16

import "fmt"

var (
	EmptyQueueErr = fmt.Errorf("cannot remove from an empty queue")
)

type Queue struct {
	Values []int
}

func New() *Queue {
	return &Queue{
		Values: make([]int, 0),
	}
}

func (q *Queue) Add(value int) {
	q.Values = append(q.Values, value)
}

func (q *Queue) Remove() (int, error) {
	if len(q.Values) < 1 {
		return 0, EmptyQueueErr
	}
	v := q.Values[0]
	q.Values = q.Values[1:]
	return v, nil
}

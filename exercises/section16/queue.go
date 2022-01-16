package section16

import "fmt"

var (
	EmptyQueueErr = fmt.Errorf("cannot remove from an empty queue")
)

type Any interface {
}

type Queue struct {
	Values []Any
}

func New() *Queue {
	return &Queue{
		Values: make([]Any, 0),
	}
}

func (q *Queue) Add(value Any) {
	q.Values = append(q.Values, value)
}

func (q *Queue) Remove() (Any, error) {
	if len(q.Values) < 1 {
		return 0, EmptyQueueErr
	}
	v := q.Values[0]
	q.Values = q.Values[1:]
	return v, nil
}

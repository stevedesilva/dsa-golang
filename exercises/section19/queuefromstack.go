package section19

import (
	"errors"

	"github.com/hashicorp/go-multierror"

	"github.com/stevedesilva/dsa-golang.git/exercises/section18"
)

type Any interface {
}

type Queue interface {
	Add(Any)
	Peek() (Any, error)
	Remove() (Any, error)
	Size() int
}

type QueueFromStack struct {
	inStack      section18.Stack
	stackAsQueue section18.Stack
}

var (
	ErrQueue = errors.New("problem with queue")
)

func New() Queue {
	return &QueueFromStack{
		inStack:      section18.New(),
		stackAsQueue: section18.New(),
	}
}

func (q *QueueFromStack) Add(item Any) {
	q.inStack.Push(item)
}

func (q *QueueFromStack) Peek() (Any, error) {
	err := q.prepForPopOrPeek()
	if err != nil {
		return nil, multierror.Append(ErrQueue, err)
	}
	item, err := q.stackAsQueue.Peek()
	if err != nil {
		return nil, multierror.Append(ErrQueue, err)
	}
	return item, nil
}

func (q *QueueFromStack) Remove() (Any, error) {
	err := q.prepForPopOrPeek()
	if err != nil {
		return nil, multierror.Append(ErrQueue, err)
	}
	item, err := q.stackAsQueue.Pop()
	if err != nil {
		return nil, multierror.Append(ErrQueue, err)
	}
	return item, nil
}

func (q *QueueFromStack) prepForPopOrPeek() error {
	for q.inStack.Size() > 0 {
		i, err := q.inStack.Pop()
		if err != nil {
			return err
		}
		q.stackAsQueue.Push(i)
	}
	return nil
}

func (q *QueueFromStack) Size() int {
	return q.inStack.Size()
}

package queue

import "fmt"

type Methods[T any] interface {
	QueuePrintJob(document T)
	Run()
	Size() int
}

type printer[T any] struct {
	q AllowedQueueFunc[T]
}

func NewPrinter[T any](documents ...T) Methods[T] {
	q1 := New[T](documents...)
	return &printer[T]{
		q: q1,
	}
}

func (p *printer[T]) QueuePrintJob(document T) {
	p.q.Enqueue(document)
}

func (p *printer[T]) Run() {
	size := p.q.Size()
	// can't use size for for loop since it keeps decreasing
	for i := 0; i < size; i++ {
		data, err := p.q.Dequeue()
		if err != nil {
			continue
		}
		p.print(data)
	}
}

func (p *printer[T]) Size() int {
	return p.q.Size()
}

func (p *printer[T]) print(val T) {
	fmt.Printf("%v \n", val)
}

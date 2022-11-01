package queue

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
	//TODO implement me
	panic("implement me")
}

func (p *printer[T]) Run() {
	//TODO implement me
	panic("implement me")
}

func (p *printer[T]) Size() int {
	return p.Size()
}

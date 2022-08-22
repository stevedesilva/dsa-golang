package scheduler

// Lazy Scheduler
// Load type without executing them

type Scheduler interface {
	Add()
	Size()
}

type Job struct {
	args []int
	f    func(...int)
}

type LazyScheduler struct {
	jobs Job
}

func (l *LazyScheduler) Size() {
	panic("implement me")
}

func (l *LazyScheduler) Add(f func(...int) int, args ...int) {

}

func New() Scheduler {
	return &LazyScheduler{}
}

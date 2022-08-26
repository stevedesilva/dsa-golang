package scheduler

// Lazy Scheduler
// Load type without executing them

type Scheduler interface {
	Add(f func(...int) int, args ...int)
	Size() int
}

type Job struct {
	args []int
	f    func(...int)
}

type LazyScheduler struct {
	jobs []Job
}

func (l *LazyScheduler) Add(f func(...int) int, args ...int) {

}

func (l *LazyScheduler) Size() int {
	return 0
}

func New() Scheduler {
	return &LazyScheduler{}
}

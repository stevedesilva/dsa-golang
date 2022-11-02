package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_printer_QueuePrintJob(t *testing.T) {
	q := NewPrinter[string]()
	q.QueuePrintJob("task1")
	q.QueuePrintJob("task2")
	q.QueuePrintJob("task3")
	q.QueuePrintJob("task4")
	assert.Equal(t, 4, q.Size())
	q.Size()
}

func Test_printer_QueuePrintJobViaNew(t *testing.T) {
	q := NewPrinter[string]("task0", "task1")
	q.QueuePrintJob("task2")
	q.QueuePrintJob("task3")
	assert.Equal(t, 4, q.Size())
	q.Size()
}

func Test_printer_QueueRunJob(t *testing.T) {
	q := NewPrinter[string]("task0", "task1")
	q.QueuePrintJob("task2")
	q.QueuePrintJob("task3")
	assert.Equal(t, 4, q.Size())
	q.Run()
	assert.Equal(t, 0, q.Size())
}

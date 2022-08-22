package scheduler_test

import (
	"testing"

	"github.com/stevedesilva/dsa-golang.git/kata/v1/scheduler"
	"github.com/stretchr/testify/assert"
)

func TestLazyScheduler_Add(t *testing.T) {
	scheduler := scheduler.New()
	scheduler.Add()
	scheduler.Add()
	want := 2
	assert.Equal(t, want, scheduler.Size)
}

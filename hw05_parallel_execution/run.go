package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) == 0 {
		return nil
	}
	if m <= 0 {
		m = 0
	}
	if n <= 0 {
		return nil
	}

	taskCh := make(chan Task)

	var wg sync.WaitGroup
	var errCount int32

	worker := func() {
		defer wg.Done()
		for task := range taskCh {
			if task() != nil && m > 0 {
				atomic.AddInt32(&errCount, 1)
			}
		}
	}

	wg.Add(n)

	for i := 0; i < n; i++ {
		go worker()
	}

	for _, task := range tasks {
		if m > 0 && atomic.LoadInt32(&errCount) >= int32(m) {
			break
		}
		taskCh <- task
	}

	close(taskCh)
	wg.Wait()

	if m > 0 && atomic.LoadInt32(&errCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

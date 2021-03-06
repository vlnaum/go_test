package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var ErrIncorrectInput = errors.New("input is incorrect")

type Task func() error

func Run(tasks []Task, n int, m int) error {
	if n <= 0 {
		return ErrIncorrectInput
	}
	tasksCh := make(chan Task)
	wg := &sync.WaitGroup{}
	allowedErrors := int32(m)

	ignoreErrors := false
	if m <= 0 {
		ignoreErrors = true
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go consumeTasks(tasksCh, wg, &allowedErrors, ignoreErrors)
	}

	for _, task := range tasks {
		if !ignoreErrors && atomic.LoadInt32(&allowedErrors) <= 0 {
			break
		}
		tasksCh <- task
	}
	close(tasksCh)
	wg.Wait()

	if allowedErrors <= 0 && !ignoreErrors {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func consumeTasks(in <-chan Task, wg *sync.WaitGroup, allowedErrors *int32, ignoreErrors bool) {
	defer wg.Done()
	var taskResult error

	for task := range in {
		if !ignoreErrors && atomic.LoadInt32(allowedErrors) <= 0 {
			return
		}

		taskResult = task()

		if !ignoreErrors && taskResult != nil {
			atomic.AddInt32(allowedErrors, -1)
		}
	}
}

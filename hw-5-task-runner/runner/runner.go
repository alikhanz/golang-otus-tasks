package runner

import (
	"errors"
)

type Runner struct {
	executedTasksCount  int
	errorsCount         int
	scheduledTasksCount int
}

type ExecutionResult struct {
	Err error
}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Run(tasks []func() error, concurrentCount int, maxErrorsCount int) error {
	resultsChan := make(chan ExecutionResult)
	finishChan := make(chan error)
	lockChan := make(chan struct{})

	defer close(finishChan)
	defer close(lockChan)
	defer close(resultsChan)

	go r.runTasks(tasks, concurrentCount, resultsChan, finishChan, lockChan)
	go r.collectResults(len(tasks), concurrentCount, maxErrorsCount, resultsChan, finishChan, lockChan)

	return <-finishChan
}

func (r *Runner) collectResults(
	tasksCount int,
	concurrentCount int,
	maxErrorsCount int,
	resultsChan <-chan ExecutionResult,
	finishChan chan<- error,
	lockChan chan<- struct{},
) {
	for {
		if r.executedTasksCount == tasksCount {
			finishChan <- nil
		}

		if r.executedTasksCount == 0 || (r.executedTasksCount%concurrentCount) == 0 {
			lockChan <- struct{}{}
		}
		select {
		case result, ok := <-resultsChan:
			if !ok {
				return
			}
			r.executedTasksCount++

			if result.Err != nil {
				r.errorsCount++
			}

			if r.errorsCount >= maxErrorsCount {
				finishChan <- errors.New("Too many errors")
				return
			}
		}
	}
}

func (r *Runner) runTasks(
	tasks []func() error,
	concurrentCount int,
	resultsChan chan ExecutionResult,
	finishChan <-chan error, lockChan <-chan struct{},
) {
	totalScheduledTasksCount := 0

	if concurrentCount > len(tasks) {
		concurrentCount = len(tasks)
	}

	for len(tasks) > totalScheduledTasksCount {
		<-lockChan
		for j := 0; j < concurrentCount; j++ {
			select {
			case <-finishChan:
				return
			default:
				go r.runTask(tasks[totalScheduledTasksCount], resultsChan)
				totalScheduledTasksCount++
			}
		}
	}
}

func (r *Runner) runTask(task func() error, resultsChan chan<- ExecutionResult) {
	resultsChan <- ExecutionResult{Err: task()}
}

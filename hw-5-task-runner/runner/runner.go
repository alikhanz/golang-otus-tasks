package runner

import (
	"errors"
	"fmt"
)

type Runner struct {
	executedTasksCount  int
	errorsCount		    int
	scheduledTasksCount int
}

type ExecutionResult struct {
	Err error
}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Run(tasks []func()error, concurrentCount int, maxErrorsCount int) error {
	resultsChan := make(chan ExecutionResult)
	finishChan := make(chan error)
	lockChan := make(chan struct{})

	defer close(finishChan)
	defer close(lockChan)
	defer close(resultsChan)

	go r.runTasks(tasks, concurrentCount, resultsChan, finishChan, lockChan)
	go r.collectResults(len(tasks), concurrentCount, maxErrorsCount, resultsChan, finishChan, lockChan)

	//return nil
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

		if r.executedTasksCount == 0 || (r.executedTasksCount % concurrentCount) == 0 {
			lockChan<- struct{}{}
		}
		fmt.Println("Select loop")
		select {
		case result, ok := <-resultsChan:
			fmt.Println("Selected results", result)
			if !ok {
				return
			}
			r.executedTasksCount++

			if result.Err != nil {
				r.errorsCount++
			}

			if r.errorsCount >= maxErrorsCount {
				fmt.Println("Too many errors")
				finishChan <- errors.New("Too many errors")
				return
			}
		}
	}
}

func (r *Runner) runTasks(tasks []func()error, concurrentCount int, resultsChan chan ExecutionResult, finishChan <-chan error, lockChan <-chan struct{}) {
	totalScheduledTasksCount := 0

	if concurrentCount > len(tasks) {
		concurrentCount = len(tasks)
	}

	for len(tasks) > totalScheduledTasksCount {
		fmt.Println("Task scheduled", totalScheduledTasksCount)
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

func (r *Runner) runTask(task func()error, resultsChan chan<- ExecutionResult) {
	resultsChan <- ExecutionResult{Err: task()}
}
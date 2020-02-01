package runner

import (
	"errors"
	"sync"
)

type ExecutionResult struct {
	Err error
}

func Run(tasks []func() error, concurrentCount int, maxErrorsCount int) error {
	if concurrentCount < 1 {
		return errors.New("concurrentCount must be more than 0")
	}

	if maxErrorsCount < 1 {
		return errors.New("maxErrorsCount must be more than 0")
	}

	resultsChan := make(chan ExecutionResult)
	exitChan := make(chan struct{}, 1)

	defer close(exitChan)

	go runTasks(tasks, concurrentCount, resultsChan, exitChan)
	err := collectResults(maxErrorsCount, resultsChan, exitChan)

	return err
}

func collectResults(
	maxErrorsCount int,
	resultsChan <-chan ExecutionResult,
	exitChan chan<- struct{},
) error {
	var executedTasksCount, errorsCount int
	var err error

	for result := range resultsChan {
		executedTasksCount++

		if result.Err != nil {
			errorsCount++
		}

		if errorsCount == maxErrorsCount {
			err = errors.New("too many errors")
			exitChan <- struct{}{}
		}
	}

	return err
}

func runTasks(
	tasks []func() error,
	concurrentCount int,
	resultsChan chan<- ExecutionResult,
	exitChan <-chan struct{},
) {
	wg := sync.WaitGroup{}
	scheduleChan := make(chan struct{}, concurrentCount)

	defer func() {
		wg.Wait()
		close(resultsChan)
		close(scheduleChan)
	}()

	scheduledTasksCount := 0

	if concurrentCount > len(tasks) {
		concurrentCount = len(tasks)
	}

	for _, task := range tasks {
		select {
			case <-exitChan:
				return
			case scheduleChan <- struct{}{}:
				wg.Add(1)
				go runTask(task, resultsChan, scheduleChan, &wg)
				scheduledTasksCount++
			}
	}
}

func runTask(
	task func() error,
	resultsChan chan<- ExecutionResult,
	scheduleChan <-chan struct{},
	wg *sync.WaitGroup,
) {
	defer func() {wg.Done()}()
	resultsChan <- ExecutionResult{Err: task()}
	<-scheduleChan
}
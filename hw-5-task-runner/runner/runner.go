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

	defer close(resultsChan)
	defer close(exitChan)

	var wg sync.WaitGroup

	go runTasks(tasks, concurrentCount, resultsChan, exitChan, &wg)
	err := collectResults(len(tasks), maxErrorsCount, resultsChan, exitChan, &wg)

	//wg.Wait()

	return err
}

func collectResults(
	tasksCount int,
	maxErrorsCount int,
	resultsChan <-chan ExecutionResult,
	exitChan chan struct{},
	wg *sync.WaitGroup,
) error {
	//defer func() { exitChan <- struct{}{} }()

	var executedTasksCount, errorsCount int
	var err error

	for {
		if executedTasksCount == tasksCount {
			exitChan <- struct{}{}
		}

		select {
		case result, ok := <-resultsChan:
			if !ok {
				return err
			}
			executedTasksCount++

			if result.Err != nil {
				errorsCount++
			}

			if errorsCount >= maxErrorsCount {
				err = errors.New("too many errors")
				exitChan <- struct{}{}
			}
		}
	}
}

func runTasks(
	tasks []func() error,
	concurrentCount int,
	resultsChan chan ExecutionResult,
	exitChan chan struct{},
	wg *sync.WaitGroup,
) {
	scheduleChan := make(chan struct{}, concurrentCount)
	defer close(scheduleChan)

	scheduledTasksCount := 0

	if concurrentCount > len(tasks) {
		concurrentCount = len(tasks)
	}

	for len(tasks) > scheduledTasksCount {
		select {
			case <-exitChan:
				return
			case scheduleChan <- struct{}{}:
				wg.Add(1)
				go runTask(tasks[scheduledTasksCount], resultsChan, scheduleChan, wg)
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
	resultsChan <- ExecutionResult{Err: task()}
	<-scheduleChan
	wg.Done()
}

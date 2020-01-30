package runner
//
//import (
//	"errors"
//	"sync"
//)
//
//type ExecutionResult struct {
//	Err error
//}
//
//func Run(tasks []func() error, concurrentCount int, maxErrorsCount int) error {
//	if concurrentCount < 1 {
//		return errors.New("concurrentCount must be more than 0")
//	}
//
//	if maxErrorsCount < 1 {
//		return errors.New("maxErrorsCount must be more than 0")
//	}
//
//	wg := &sync.WaitGroup{}
//	i := 0
//
//	scheduleChan := make(chan struct{}, concurrentCount)
//	resultsChan := make(chan ExecutionResult)
//
//	for len(tasks) < i {
//		scheduleChan <- struct{}{}
//		go runTask(tasks[i], wg, resultsChan, scheduleChan)
//		i++
//	}
//
//
//
//	wg.Wait()
//
//	return nil
//}
//
//func runTask(
//	task func() error,
//	wg *sync.WaitGroup,
//	resultsChan chan<- ExecutionResult,
//	scheduleChan <-chan struct{},
//) {
//	result := task()
//	wg.Done()
//	resultsChan <- ExecutionResult{Err: result}
//	<-scheduleChan
//}
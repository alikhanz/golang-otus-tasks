package runner

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const TasksCount = 20
const ConcurrentCount = 5
const ErrorsCount = 2

func TestRunner_RunSuccessOneTask(t *testing.T) {
	tasks := make([]func() error, 0, 1)

	tasks = append(tasks, func() error {
		time.Sleep(1*time.Second)
		return nil
	})

	err := Run(tasks, 1, 1)

	assert.NoError(t, err)
}

func TestRunner_RunErrors(t *testing.T) {
	tt := []struct {
		concurrentCount int
		errorsCount int
	}{
		{
			concurrentCount: 0,
			errorsCount:     1,
		},
		{
			concurrentCount: 1,
			errorsCount:     0,
		},
		{
			concurrentCount: -1,
			errorsCount:     1,
		},
		{
			concurrentCount: 1,
			errorsCount:     -1,
		},
	}

	tasks := make([]func() error, 0, 1)

	tasks = append(tasks, func() error {
		return nil
	})

	for _, v := range tt {
		err := Run(tasks, v.concurrentCount, v.errorsCount)
		assert.Error(t, err)
	}
}

func TestRunner_RunFailedOneTask(t *testing.T) {
	tasks := make([]func() error, 0, 1)

	tasks = append(tasks, func() error {
		return errors.New("Mew mew")
	})

	err := Run(tasks, 1, 1)

	assert.Error(t, err)
}

func TestRunner_RunSuccessMany(t *testing.T) {
	tasks := make([]func() error, 0, 1)
	var executedCount int

	for i := 0; i < TasksCount; i++ {
		tasks = append(tasks, func() error {
			executedCount++
			return nil
		})
	}

	tasks = append(tasks, func() error {
		time.Sleep(1 * time.Second)
		return nil
	})
	err := Run(tasks, ConcurrentCount, ErrorsCount)

	assert.NoError(t, err)
	assert.Equal(t, executedCount, TasksCount)
}

func TestRunner_RunFailMany(t *testing.T) {
	tasks := make([]func() error, 0, 1)
	var executedCount int

	for i := 0; i < TasksCount; i++ {
		tasks = append(tasks, func() error {
			executedCount++
			return errors.New("Mew mew")
		})
	}

	err := Run(tasks, ConcurrentCount, ErrorsCount)

	assert.Error(t, err)
	assert.True(t, executedCount < ConcurrentCount+ErrorsCount)
}

package runner

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TasksCount = 20
const ConcurrentCount = 5
const ErrorsCount = 2

func TestRunner_RunSuccessOneTask(t *testing.T) {
	tasks := make([]func() error, 0, 1)

	tasks = append(tasks, func() error {
		return nil
	})

	runner := New()
	err := runner.Run(tasks, 1, 1)

	assert.NoError(t, err)
}

func TestRunner_RunFailedOneTask(t *testing.T) {
	tasks := make([]func() error, 0, 1)

	tasks = append(tasks, func() error {
		return errors.New("Mew mew")
	})

	runner := New()
	err := runner.Run(tasks, 1, 1)

	assert.Error(t, err)
}

func TestRunner_RunSuccessMany(t *testing.T) {
	tasks := make([]func() error, 0, 1)

	for i := 0; i < TasksCount; i++ {
		tasks = append(tasks, func() error {
			return nil
		})
	}

	runner := New()
	err := runner.Run(tasks, ConcurrentCount, ErrorsCount)

	assert.NoError(t, err)
	assert.Equal(t, runner.executedTasksCount, TasksCount)
}

func TestRunner_RunFailMany(t *testing.T) {
	tasks := make([]func() error, 0, 1)

	for i := 0; i < TasksCount; i++ {
		tasks = append(tasks, func() error {
			return errors.New("Mew mew")
		})
	}

	runner := New()
	err := runner.Run(tasks, ConcurrentCount, ErrorsCount)

	assert.Error(t, err)
	assert.True(t, runner.executedTasksCount < ConcurrentCount+ErrorsCount)
}

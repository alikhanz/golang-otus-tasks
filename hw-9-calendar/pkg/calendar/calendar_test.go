package calendar

import (
	"github.com/alikhanz/golang-otus-tasks/hw-9-calendar/pkg/event"
	"github.com/alikhanz/golang-otus-tasks/hw-9-calendar/pkg/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalendar_CreateEvent(t *testing.T) {
	c := makeCalendar()
	id, err := c.CreateEvent("Test title", "Test description", time.Now(), false)
	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.Nil)
}

func TestCalendar_UpdateEvent(t *testing.T) {
	c := makeCalendar()
	id, err := c.CreateEvent("Test title", "Test description", time.Now(), false)
	assert.NoError(t, err)
	assert.NotEqual(t, id, uuid.Nil)

	e, err := c.FetchById(id)

	assert.NoError(t, err)
	assert.NotEmpty(t, e)

	e.Title = "Updated title"
	e.Description = "Updated description"

	err = c.UpdateEvent(e)
	assert.NoError(t, err)

	e, err = c.FetchById(e.Id)
	assert.NoError(t, err)
	assert.Equal(t, e.Title, "Updated title")
	assert.Equal(t, e.Description, "Updated description")
}

func TestCalendar_UpdateEventError(t *testing.T) {
	c := makeCalendar()
	e := event.NewEvent("title", "description", time.Now(), false)
	err := c.UpdateEvent(e)
	assert.Error(t, err)
	assert.IsType(t, err, NewEventUpdateError{})
}

func makeCalendar() *Calendar {
	return NewCalendar(storage.NewMemoryStorage())
}
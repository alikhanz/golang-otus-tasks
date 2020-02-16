package memory

import (
	"github.com/alikhanz/golang-otus-tasks/hw-8-calendar/pkg/event"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMemory_SaveNewEvent(t *testing.T) {
	s := NewMemoryStorage()
	e := makeEvent()

	err := s.Save(&e)

	assert.NoError(t, err, "Failed saving new event")
	assert.NotEqual(t, e.Id(), uuid.Nil)
}

func TestMemory_UpdateEvent(t *testing.T) {
	s := NewMemoryStorage()
	e := makeEvent()

	err := s.Save(&e)
	assert.NoError(t, err, "Failed saving new event")

	e.SetTitle("Updated title")
	e.SetDescription("Updated description")
	e.SetDateTime(time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local))
	e.SetRepeatable(true)

	err = s.Save(&e)
	assert.NoError(t, err, "Failed event save")

	sEvent, err := s.FetchById(e.Id())
	assert.NoError(t, err, "Failed fetch event from storage")

	assert.NotEqual(t, e.Id(), uuid.Nil)
	assert.Equal(t, e, *sEvent)
}

func TestMemory_RemoveEvent(t *testing.T) {
	s := NewMemoryStorage()
	e := makeEvent()

	err := s.Save(&e)

	assert.NoError(t, err, "Failed saving new event")
	assert.NotEqual(t, e.Id(), uuid.Nil)

	err = s.Remove(e)
	assert.NoError(t, err, "Failed removing event")

	_, err = s.FetchById(e.Id())
	assert.Error(t, err)
	assert.IsType(t, err, event.NotFoundError{})
}

func TestMemory_RemoveNotExistEvent(t *testing.T) {
	s := NewMemoryStorage()
	e := makeEvent()
	e.SetId(uuid.New())

	err := s.Remove(e)
	assert.Error(t, err)
	assert.IsType(t, err, event.NotFoundError{})
}

func TestMemory_FetchById(t *testing.T) {
	s := NewMemoryStorage()
	e := makeEvent()

	err := s.Save(&e)

	assert.NoError(t, err, "Failed saving new event")
	assert.NotEqual(t, e.Id(), uuid.Nil)

	ev, err := s.FetchById(e.Id())
	assert.NotNil(t, ev)
	assert.NoError(t, err)
}

func TestMemory_FetchByIdNotExist(t *testing.T) {
	s := NewMemoryStorage()
	ev, err := s.FetchById(uuid.New())
	assert.Nil(t, ev)
	assert.Error(t, err)
	assert.IsType(t, err, event.NotFoundError{})
}

func TestMemory_FetchBetweenDates(t *testing.T) {
	s := NewMemoryStorage()

	for i := 0; i < 3; i++ {
		e := makeEvent()
		e.SetDateTime(time.Date(2000, time.Month(1+i), 1, 0, 0, 0, 0, time.Local))
		err := s.Save(&e)
		assert.NoError(t, err, "Failed event save")
	}

	from := time.Date(2000, 1, 31, 0, 0, 0, 0, time.Local)
	to := time.Date(2000, 2, 1, 0, 0, 0, 0, time.Local)
	events, err := s.FetchBetweenDates(from, to)
	assert.NoError(t, err)
	assert.Len(t, events, 1)

	dt := time.Date(2000, 2, 1, 0, 0, 0, 0, time.Local)
	assert.Equal(t, events[0].DateTime(), dt)
}

func makeEvent() event.Event {
	return *event.NewEvent("Test", "Test description", time.Now(), false)
}

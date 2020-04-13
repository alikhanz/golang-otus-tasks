package calendar

import (
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/event"
	"github.com/google/uuid"
	"time"
)

type Storage interface {
	Remove(event event.Event) error
	Save(event event.Event) (uuid.UUID, error)
	FetchById(uuid uuid.UUID) (event.Event, error)
	FetchBetweenDates(from, to time.Time) ([]event.Event, error)
	All() ([]event.Event, error)
}
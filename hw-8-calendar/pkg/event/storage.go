package event

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type NotFoundError struct {
	Event Event
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("event not found in storage %v", n.Event)
}

type Storage interface {
	Remove(event Event) error
	Save(event *Event) error
	FetchById(uuid uuid.UUID) (*Event, error)
	FetchBetweenDates(from, to time.Time) ([]Event, error)
	All() ([]Event, error)
}
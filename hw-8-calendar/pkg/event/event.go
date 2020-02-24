package event

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	Id          uuid.UUID
	Title       string
	Description string
	DateTime    time.Time
	Repeatable  bool
}

func NewEvent(
	title string,
	description string,
	dateTime time.Time,
	repeatable bool,
) Event {
	return Event{
		Title:       title,
		Description: description,
		DateTime:    dateTime,
		Repeatable:  repeatable,
	}
}

func (e *Event) IsNew() bool {
	return e.Id == uuid.Nil
}
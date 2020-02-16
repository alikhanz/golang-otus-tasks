package event

import (
	"github.com/google/uuid"
	"time"
)

type Event struct {
	id			uuid.UUID
	title       string
	description string
	dateTime    time.Time
	repeatable  bool
}

func NewEvent(
	title string,
	description string,
	dateTime time.Time,
	repeatable bool,
) *Event {
	return &Event{
		title:       title,
		description: description,
		dateTime:    dateTime,
		repeatable:  repeatable,
	}
}

func (e *Event) Id() uuid.UUID {
	return e.id
}

func (e *Event) SetId(id uuid.UUID) {
	e.id = id
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) {
	e.title = title
}

func (e *Event) Description() string {
	return e.description
}

func (e *Event) SetDescription(description string) {
	e.description = description
}

func (e Event) DateTime() time.Time {
	return e.dateTime
}

func (e *Event) SetDateTime(dateTime time.Time) {
	e.dateTime = dateTime
}

func (e *Event) Repeatable() bool {
	return e.repeatable
}

func (e *Event) SetRepeatable(repeatable bool) {
	e.repeatable = repeatable
}

func (e *Event) IsNew() bool {
	return e.id == uuid.Nil
}
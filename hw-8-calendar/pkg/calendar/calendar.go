package calendar

import (
	"fmt"
	"github.com/alikhanz/golang-otus-tasks/hw-8-calendar/pkg/event"
	"github.com/google/uuid"
	"time"
)

type NewEventUpdateError struct {

}

type UpdateEventError struct {
	Err error
}

func (n NewEventUpdateError) Error() string {
	return "Cannot update new event"
}

func (u UpdateEventError) Error() string {
	return fmt.Sprintf("Event update error: %s", u.Err)
}

type Calendar struct {
	eventStorage Storage
}

func NewCalendar(storage Storage) *Calendar {
	return &Calendar{eventStorage: storage}
}

func (c *Calendar) CreateEvent(
	title string,
	description string,
	dateTime time.Time,
	repeatable bool,
) (uuid.UUID, error) {
	e := event.NewEvent(
		title,
		description,
		dateTime,
		repeatable,
	)

	id, err := c.eventStorage.Save(e)

	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (c *Calendar) UpdateEvent(e event.Event) error {
	if e.IsNew() {
		return NewEventUpdateError{}
	}

	_, err := c.eventStorage.Save(e)

	if err != nil {
		return UpdateEventError{Err: err}
	}

	return nil
}

func (c *Calendar) FetchById(uuid uuid.UUID) (event.Event, error) {
	return c.eventStorage.FetchById(uuid)
}

func (c *Calendar) FetchBetweenDates(from, to time.Time) ([]event.Event, error) {
	return c.eventStorage.FetchBetweenDates(from, to)
}

func (c *Calendar) RemoveEvent(e event.Event) error {
	return c.eventStorage.Remove(e)
}
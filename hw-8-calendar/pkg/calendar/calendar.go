package calendar

import "github.com/alikhanz/golang-otus-tasks/hw-8-calendar/pkg/event"

type Calendar struct {
	eventStorage event.Storage
}

func (c Calendar) EventStorage() event.Storage {
	return c.eventStorage
}

func NewCalendar(storage event.Storage) *Calendar {
	return &Calendar{eventStorage: storage}
}
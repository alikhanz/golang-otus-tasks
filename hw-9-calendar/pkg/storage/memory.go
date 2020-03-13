package storage

import (
	"github.com/alikhanz/golang-otus-tasks/hw-9-calendar/pkg/event"
	"github.com/google/uuid"
	"time"
)

type Memory struct {
	events map[uuid.UUID]event.Event
}

func NewMemoryStorage() *Memory {
	return &Memory{events: make(map[uuid.UUID]event.Event)}
}

func (m *Memory) FetchById(uuid uuid.UUID) (event.Event, error) {
	e, ok := m.events[uuid]

	if !ok {
		return event.Event{}, event.NotFoundError{Event: e}
	}

	return e, nil
}

func (m *Memory) FetchBetweenDates(from, to time.Time) ([]event.Event, error) {
	result := make([]event.Event, 0)

	 for _, ev := range m.events {
	 	if ev.DateTime.Unix() >= from.Unix() && ev.DateTime.Unix() <= to.Unix() {
			result = append(result, ev)
		}
	 }

	 return result, nil
}

func (m *Memory) Remove(e event.Event) error {
	_, ok := m.events[e.Id]

	if !ok {
		return event.NotFoundError{Event: e}
	}

	delete(m.events, e.Id)
	return nil
}

//Save save event in memory.
func (m *Memory) Save(e event.Event) (uuid.UUID, error) {
	if e.IsNew() {
		e.Id = uuid.New()
	}

	m.events[e.Id] = e
	return e.Id, nil

}

//All unordered events list
func (m *Memory) All() ([]event.Event, error) {
	res := make([]event.Event, 0, len(m.events))

	for _, v := range m.events {
		res = append(res, v)
	}

	return res, nil
}
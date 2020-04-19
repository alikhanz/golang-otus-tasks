package mapper

import (
	calendarPb "github.com/alikhanz/golang-otus-tasks/calendar/api_pb/api/protobuf"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/event"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type EventMapper struct {
}

func (eventMapper *EventMapper) MapPbToModel(inEvent *calendarPb.Event) (event.Event, error) {
	t, err := ptypes.Timestamp(inEvent.Time)

	if err != nil {
		return event.Event{}, errors.Wrap(err, "pb event to model mapping failed")
	}

	resEvent := event.NewEvent(
		inEvent.Title,
		inEvent.Description,
		t,
		inEvent.Repeatable,
	)

	// При ошибке просто получим nil uuid
	id, _ := uuid.Parse(inEvent.EventId)
	resEvent.Id = id

	return resEvent, nil
}

func (eventMapper *EventMapper) MapModelToPb(ev event.Event) *calendarPb.Event {
	t, _ := ptypes.TimestampProto(ev.DateTime)

	return &calendarPb.Event{
		EventId:     ev.Id.String(),
		Title:       ev.Title,
		Description: ev.Description,
		Time:        t,
		Repeatable:  ev.Repeatable,
	}
}

func (eventMapper *EventMapper) MapModelListToPb(events []event.Event) []*calendarPb.Event {
	result := make([]*calendarPb.Event, 0, len(events))

	for _, ev := range events {
		result = append(result, eventMapper.MapModelToPb(ev))
	}

	return result
}

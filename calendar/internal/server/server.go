package server

import (
	"context"
	"github.com/alikhanz/golang-otus-tasks/calendar/api_pb/api/protobuf"
	"github.com/alikhanz/golang-otus-tasks/calendar/internal/mapper"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/calendar"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Run() {

}

var eventMapper *mapper.EventMapper

type CalendarServer struct {
	cal calendar.Calendar
}

func (c CalendarServer) CreateEvent(
	ctx context.Context,
	request *calendarPb.CreateEventRequest,
) (*calendarPb.Event, error) {
	err := ValidateCreateEventRequest(request)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	t, err := ptypes.Timestamp(request.Time)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid date")
	}

	eId, err := c.cal.CreateEvent(
		request.Title,
		request.Description,
		t,
		request.Repeatable,
	)

	if err != nil {
		return nil, status.Error(codes.Internal, "Cannot save event")
	}

	ev, err := c.cal.FetchById(eId)

	if err != nil {
		return nil, status.Error(codes.Internal, "Event fetch failed")
	}

	return eventMapper.MapModelToPb(ev), nil
}

func (c CalendarServer) UpdateEvent(
	ctx context.Context,
	event *calendarPb.Event,
) (*calendarPb.Event, error) {
	panic("implement me")
}

func (c CalendarServer) RemoveEvent(
	ctx context.Context,
	request *calendarPb.RemoveEventRequest,
) (*calendarPb.RemoveEventResponse, error) {
	panic("implement me")
}

func (c CalendarServer) GetEventsList(
	ctx context.Context,
	request *calendarPb.GetEventsListRequest,
) (*calendarPb.EventsList, error) {
	panic("implement me")
}

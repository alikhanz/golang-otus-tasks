package server

import (
	calendarPb "github.com/alikhanz/golang-otus-tasks/calendar/api_pb/api/protobuf"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/calendar"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/validator"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Validator struct {
	cal *calendar.Calendar
}

func NewValidator(cal *calendar.Calendar) *Validator {
	return &Validator{cal: cal}
}

func (*Validator) ValidateCreateEventRequest(r *calendarPb.CreateEventRequest) error {
	v := validator.NewValidator()
	v.ValidateString(r.Title, &validator.NotEmptyStringRule{Field: "Title"})
	v.ValidateString(r.Description, &validator.NotEmptyStringRule{Field: "Description"})

	if v.HasErrors() {
		return errors.New(v.RenderErrors())
	}

	return nil
}

func (val *Validator) ValidateEventExist(e *calendarPb.Event) error {
	v := validator.NewValidator()
	id, err := uuid.Parse(e.EventId)

	if err != nil {
		v.AddError("uuid parse failed")
	}

	_, err = val.cal.FetchById(id)

	if err != nil {
		v.AddError("event not found")
	}
}

func (*Validator) ValidateEventFields(e *calendarPb.Event) error {
	v := validator.NewValidator()
	v.ValidateString(e.Title, &validator.NotEmptyStringRule{Field: "Title"})
	v.ValidateString(e.Description, &validator.NotEmptyStringRule{Field: "Description"})

	_, err := ptypes.Timestamp(e.Time)

	if err != nil {
		v.AddError("Time has incorrect date")
	}

	if v.HasErrors() {
		return errors.New(v.RenderErrors())
	}

	return nil

}

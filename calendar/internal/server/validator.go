package server

import (
	calendarPb "github.com/alikhanz/golang-otus-tasks/calendar/api_pb/api/protobuf"
	"github.com/alikhanz/golang-otus-tasks/calendar/pkg/validator"
	"github.com/pkg/errors"
)

func ValidateCreateEventRequest(r *calendarPb.CreateEventRequest) error {
	v := validator.NewValidator()
	v.ValidateString(r.Title, &validator.NotEmptyStringRule{Field: "Title"})
	v.ValidateString(r.Description, &validator.NotEmptyStringRule{Field: "Description"})

	if v.HasErrors() {
		return errors.New(v.RenderErrors())
	}

	return nil
}
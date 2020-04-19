package event

import (
	"fmt"
	"github.com/google/uuid"
)

type NotFoundError struct {
	Id uuid.UUID
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("event %s not found in storage", n.Id.String())
}

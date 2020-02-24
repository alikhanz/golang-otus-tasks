package event

import (
	"fmt"
)

type NotFoundError struct {
	Event Event
}

func (n NotFoundError) Error() string {
	return fmt.Sprintf("event not found in storage %v", n.Event)
}
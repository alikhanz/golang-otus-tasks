package validator

import (
	"strings"
)

type Rule interface {
	ErrorMessage() string
}

type StringRule interface {
	Rule
	Validate(string) bool
}

type NotEmptyStringRule struct {
	Field string
}

func (r *NotEmptyStringRule) ErrorMessage() string {
	b := strings.Builder{}

	if len(r.Field) > 0 {
		b.WriteString(r.Field + " ")
	}
	b.WriteString("value cannot be empty string")

	return b.String()
}

func (r *NotEmptyStringRule) Validate(value string) bool {
	return len(value) > 0
}

package validator

import "strings"

type Errors struct {
	errors []string
}

func (e *Errors) Add(err string) {
	e.errors = append(e.errors, err)
}

func (e *Errors) Render() string {
	return strings.Join(e.errors, "\n")
}

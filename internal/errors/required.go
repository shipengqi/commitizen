package errors

import "fmt"

type RequiredErr struct {
	field string
}

func (e RequiredErr) Error() string {
	return fmt.Sprintf("%s is required", e.field)
}

func NewRequiredErr(field string) error {
	return RequiredErr{field: field}
}

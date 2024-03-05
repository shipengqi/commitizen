package render

import (
	"errors"
	"fmt"
)

var (
	ErrCanceled = errors.New("canceled")
)

type MissingErr struct {
	field string
}

func (e MissingErr) Error() string {
	return fmt.Sprintf("%s is required", e.field)
}

func NewMissingErr(field string) error {
	return MissingErr{field: field}
}

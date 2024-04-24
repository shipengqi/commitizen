package render

import (
	"fmt"
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

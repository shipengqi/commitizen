package errors

import "fmt"

type MissingErr struct {
	name  string
	field string
}

func (e MissingErr) Error() string {
	if e.name == "" {
		return fmt.Sprintf("missing required field `%s`", e.field)
	}
	return fmt.Sprintf("'%s' missing required field: %s", e.name, e.field)
}

func NewMissingErr(field string, name ...string) error {
	err := MissingErr{field: field}
	if len(name) > 0 {
		err.name = name[0]
	}
	return err
}

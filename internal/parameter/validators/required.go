package validators

import (
	"fmt"
	"strings"
)

func Required(name string, trim bool) func(string) error {
	return func(str string) error {
		if trim {
			str = strings.TrimSpace(str)
		}
		if len(str) == 0 {
			return fmt.Errorf("'%s' cannot be empty", name)
		}
		return nil
	}
}

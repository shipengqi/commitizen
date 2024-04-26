package validators

import (
	"fmt"
	"strings"

	"github.com/shipengqi/commitizen/internal/errors"
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

func MultiRequired(name string) func([]string) error {
	return func(strs []string) error {
		if len(strs) == 0 {
			return errors.NewRequiredErr(name)
		}
		return nil
	}
}

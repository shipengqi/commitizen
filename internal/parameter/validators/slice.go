package validators

import "github.com/shipengqi/commitizen/internal/errorsx"

func MultiRequired(name string) func([]string) error {
	return func(vals []string) error {
		if len(vals) == 0 {
			return errorsx.NewRequiredErr(name)
		}
		return nil
	}
}

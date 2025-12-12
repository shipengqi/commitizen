package helpers

import "github.com/shipengqi/commitizen/internal/errorsx"

func GetValueFromYAML[T any](data map[string]any, key string) (T, error) {
	var (
		res T
		ok  bool
		v   any
	)

	v, ok = data[key]
	if !ok {
		return res, errorsx.NewMissingErr(key)
	}
	res, ok = v.(T)
	if !ok {
		return res, errorsx.ErrType
	}
	return res, nil
}

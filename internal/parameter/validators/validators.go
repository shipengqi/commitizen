package validators

type Validator[T string | []string] func(T) error

func Group[T string | []string](validators ...Validator[T]) Validator[T] {
	return func(t T) error {
		for _, validator := range validators {
			if err := validator(t); err != nil {
				return err
			}
		}
		return nil
	}
}

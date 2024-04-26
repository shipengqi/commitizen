package validators

type Validator func(string) error

func Group(validators ...Validator) Validator {
	return func(str string) error {
		for _, validator := range validators {
			if err := validator(str); err != nil {
				return err
			}
		}
		return nil
	}
}

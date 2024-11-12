package validators

import (
	"errors"
	"fmt"
	"strconv"
)

func Max(maxVal int) func(string) error {
	return func(str string) error {
		v, err := strconv.Atoi(str)
		if err != nil {
			return errors.New("invalid integer")
		}
		if v > maxVal {
			return fmt.Errorf("value must less than or equal to %d", maxVal)
		}
		return nil
	}
}

func Min(minVal int) func(string) error {
	return func(str string) error {
		v, err := strconv.Atoi(str)
		if err != nil {
			return errors.New("invalid integer")
		}
		if v < minVal {
			return fmt.Errorf("value must less than or equal to %d", minVal)
		}
		return nil
	}
}

package validators

import (
	"errors"
	"fmt"
	"strconv"
)

func Max(max int) func(string) error {
	return func(str string) error {
		v, err := strconv.Atoi(str)
		if err != nil {
			return errors.New("invalid integer")
		}
		if v > max {
			return fmt.Errorf("value must less than or equal to %d", max)
		}
		return nil
	}
}

func Min(min int) func(string) error {
	return func(str string) error {
		v, err := strconv.Atoi(str)
		if err != nil {
			return errors.New("invalid integer")
		}
		if v < min {
			return fmt.Errorf("value must less than or equal to %d", min)
		}
		return nil
	}
}

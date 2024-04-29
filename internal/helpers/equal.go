package helpers

import (
	"bytes"
	"errors"
	"reflect"
)

// Equal asserts that two objects are equal.
//
//	helpers.Equal(123, 123)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses). Function equality
// cannot be determined and will always fail.
func Equal(expected, actual interface{}) bool {
	if err := validateEqualArgs(expected, actual); err != nil {
		// invalid operation
		return false
	}
	return ObjectsAreEqual(expected, actual)
}

// NotEqual asserts that the specified values are NOT equal.
//
//	helpers.NotEqual(obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
func NotEqual(expected, actual interface{}, msgAndArgs ...interface{}) bool {
	if err := validateEqualArgs(expected, actual); err != nil {
		// invalid operation
		return false
	}

	return !ObjectsAreEqual(expected, actual)
}

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}

// validateEqualArgs checks whether provided arguments can be safely used in the
// Equal/NotEqual functions.
func validateEqualArgs(expected, actual interface{}) error {
	if expected == nil && actual == nil {
		return nil
	}

	if isFunction(expected) || isFunction(actual) {
		return errors.New("cannot take func type as argument")
	}
	return nil
}

func isFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}

package helpers

import (
	"reflect"
	"strings"
)

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
//	helpers.Contains("Hello World", "World")
//	helpers.Contains(["Hello", "World"], "World")
//	helpers.Contains({"Hello": "World"}, "Hello")
func Contains(s, contains interface{}, msgAndArgs ...interface{}) bool {
	ok, found := containsElement(s, contains)
	if !ok {
		return false
	}

	return found
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
//	helpers.NotContains("Hello World", "Earth")
//	helpers.NotContains(["Hello", "World"], "Earth")
//	helpers.NotContains({"Hello": "World"}, "Earth")
func NotContains(s, contains interface{}, msgAndArgs ...interface{}) bool {
	ok, found := containsElement(s, contains)
	if !ok {
		return false
	}

	return !found
}

// containsElement try loop over the list check if the list includes the element.
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func containsElement(list interface{}, element interface{}) (ok, found bool) {

	listValue := reflect.ValueOf(list)
	listType := reflect.TypeOf(list)
	if listType == nil {
		return false, false
	}
	listKind := listType.Kind()
	defer func() {
		if e := recover(); e != nil {
			ok = false
			found = false
		}
	}()

	if listKind == reflect.String {
		elementValue := reflect.ValueOf(element)
		return true, strings.Contains(listValue.String(), elementValue.String())
	}

	if listKind == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if ObjectsAreEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if ObjectsAreEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false

}

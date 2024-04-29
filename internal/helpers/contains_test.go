package helpers

import (
	"fmt"
	"testing"
)

func TestContainsNotContains(t *testing.T) {

	type A struct {
		Name, Value string
	}
	list := []string{"Foo", "Bar"}

	complexList := []*A{
		{"b", "c"},
		{"d", "e"},
		{"g", "h"},
		{"j", "k"},
	}
	simpleMap := map[interface{}]interface{}{"Foo": "Bar"}
	var zeroMap map[interface{}]interface{}

	cases := []struct {
		expected interface{}
		actual   interface{}
		result   bool
	}{
		{"Hello World", "Hello", true},
		{"Hello World", "Salut", false},
		{list, "Bar", true},
		{list, "Salut", false},
		{complexList, &A{"g", "h"}, true},
		{complexList, &A{"g", "e"}, false},
		{simpleMap, "Foo", true},
		{simpleMap, "Bar", false},
		{zeroMap, "Bar", false},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Contains(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := Contains(c.expected, c.actual)

			if res != c.result {
				if res {
					t.Errorf("Contains(%#v, %#v) should return true:\n\t%#v contains %#v", c.expected, c.actual, c.expected, c.actual)
				} else {
					t.Errorf("Contains(%#v, %#v) should return false:\n\t%#v does not contain %#v", c.expected, c.actual, c.expected, c.actual)
				}
			}
		})
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("NotContains(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := NotContains(c.expected, c.actual)

			// NotContains should be inverse of Contains. If it's not, something is wrong
			if res == Contains(c.expected, c.actual) {
				if res {
					t.Errorf("NotContains(%#v, %#v) should return true:\n\t%#v does not contains %#v", c.expected, c.actual, c.expected, c.actual)
				} else {
					t.Errorf("NotContains(%#v, %#v) should return false:\n\t%#v contains %#v", c.expected, c.actual, c.expected, c.actual)
				}
			}
		})
	}
}

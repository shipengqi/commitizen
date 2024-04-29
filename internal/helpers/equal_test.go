package helpers

import (
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	type myType string

	var m map[string]interface{}

	cases := []struct {
		expected interface{}
		actual   interface{}
		result   bool
		remark   string
	}{
		{"Hello World", "Hello World", true, ""},
		{123, 123, true, ""},
		{123.5, 123.5, true, ""},
		{[]byte("Hello World"), []byte("Hello World"), true, ""},
		{nil, nil, true, ""},
		{int32(123), int32(123), true, ""},
		{uint64(123), uint64(123), true, ""},
		{myType("1"), myType("1"), true, ""},
		{&struct{}{}, &struct{}{}, true, "pointer equality is based on equality of underlying value"},
		{[]string{"str1", "str2"}, []string{"str1", "str2"}, true, ""},
		{[]int{1, 2}, []int{1, 2}, true, ""},
		{true, true, true, ""},

		// Not expected to be equal
		{m["bar"], "something", false, ""},
		{myType("1"), myType("2"), false, ""},

		// A case that might be confusing, especially with numeric literals
		{10, uint(10), false, ""},

		{[]string{"str1", "str2"}, []string{"str1", "str3"}, false, ""},
		{[]int{1, 2}, []int{1, 3}, false, ""},
		{true, false, false, ""},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := Equal(c.expected, c.actual)

			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	type myStructType struct{}

	cases := []struct {
		expected interface{}
		actual   interface{}
		result   bool
	}{
		// cases that are expected not to match
		{"Hello World", "Hello World!", true},
		{123, 1234, true},
		{123.5, 123.55, true},
		{[]byte("Hello World"), []byte("Hello World!"), true},
		{nil, new(myStructType), true},

		// cases that are expected to match
		{nil, nil, false},
		{"Hello World", "Hello World", false},
		{123, 123, false},
		{123.5, 123.5, false},
		{[]byte("Hello World"), []byte("Hello World"), false},
		{new(myStructType), new(myStructType), false},
		{&struct{}{}, &struct{}{}, false},
		{func() int { return 23 }, func() int { return 24 }, false},
		// A case that might be confusing, especially with numeric literals
		{int(10), uint(10), true},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("NotEqual(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := NotEqual(c.expected, c.actual)

			if res != c.result {
				t.Errorf("NotEqual(%#v, %#v) should return %#v", c.expected, c.actual, c.result)
			}
		})
	}
}

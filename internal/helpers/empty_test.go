package helpers

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {

	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	var tiP *time.Time
	var tiNP time.Time
	var s *string
	var f *os.File
	sP := &s
	x := 1
	xP := &x

	type TString string
	type TStruct struct {
		x int
	}

	assert.True(t, Empty(""), "Empty string is empty")
	assert.True(t, Empty(nil), "Nil is empty")
	assert.True(t, Empty([]string{}), "Empty string array is empty")
	assert.True(t, Empty(0), "Zero int value is empty")
	assert.True(t, Empty(false), "False value is empty")
	assert.True(t, Empty(make(chan struct{})), "Channel without values is empty")
	assert.True(t, Empty(s), "Nil string pointer is empty")
	assert.True(t, Empty(f), "Nil os.File pointer is empty")
	assert.True(t, Empty(tiP), "Nil time.Time pointer is empty")
	assert.True(t, Empty(tiNP), "time.Time is empty")
	assert.True(t, Empty(TStruct{}), "struct with zero values is empty")
	assert.True(t, Empty(TString("")), "empty aliased string is empty")
	assert.True(t, Empty(sP), "ptr to nil value is empty")
	assert.True(t, Empty([1]int{}), "array is state")

	assert.False(t, Empty("something"), "Non Empty string is not empty")
	assert.False(t, Empty(errors.New("something")), "Non nil object is not empty")
	assert.False(t, Empty([]string{"something"}), "Non empty string array is not empty")
	assert.False(t, Empty(1), "Non-zero int value is not empty")
	assert.False(t, Empty(true), "True value is not empty")
	assert.False(t, Empty(chWithValue), "Channel with values is not empty")
	assert.False(t, Empty(TStruct{x: 1}), "struct with initialized values is empty")
	assert.False(t, Empty(TString("abc")), "non-empty aliased string is empty")
	assert.False(t, Empty(xP), "ptr to non-nil value is not empty")
	assert.False(t, Empty([1]int{42}), "array is not state")
}

func TestNotEmpty(t *testing.T) {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}

	assert.False(t, NotEmpty(""), "Empty string is empty")
	assert.False(t, NotEmpty(nil), "Nil is empty")
	assert.False(t, NotEmpty([]string{}), "Empty string array is empty")
	assert.False(t, NotEmpty(0), "Zero int value is empty")
	assert.False(t, NotEmpty(false), "False value is empty")
	assert.False(t, NotEmpty(make(chan struct{})), "Channel without values is empty")
	assert.False(t, NotEmpty([1]int{}), "array is state")

	assert.True(t, NotEmpty("something"), "Non Empty string is not empty")
	assert.True(t, NotEmpty(errors.New("something")), "Non nil object is not empty")
	assert.True(t, NotEmpty([]string{"something"}), "Non empty string array is not empty")
	assert.True(t, NotEmpty(1), "Non-zero int value is not empty")
	assert.True(t, NotEmpty(true), "True value is not empty")
	assert.True(t, NotEmpty(chWithValue), "Channel with values is not empty")
	assert.True(t, NotEmpty([1]int{42}), "array is not state")
}

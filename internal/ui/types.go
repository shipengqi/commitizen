package ui

import (
	"errors"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

// EchoMode sets the input behavior of the text input field.
// EchoMode is an alias for the textinput.EchoMode.
type EchoMode int

const (
	// EchoNormal displays text as is. This is the default behavior.
	EchoNormal EchoMode = iota

	// EchoPassword displays the EchoCharacter mask instead of actual
	// characters. This is commonly used for password fields.
	EchoPassword

	// EchoNone displays nothing as characters are entered. This is commonly
	// seen for password fields on the command line.
	EchoNone
)

const (
	DefaultValidateOkPrefix  = "✔"
	DefaultValidateErrPrefix = "✘"
	DefaultTextAreaMaxHeight = 5
	DefaultTextAreaHeight    = 5
	DefaultSelectWidth       = 20
	DefaultSelectHeight      = 12

	ColorPrompt      = "2"
	colorValidateOk  = "2"
	colorValidateErr = "1"
)

const DONE = "DONE"

// DefaultValidateFunc is a verification function that does nothing
func DefaultValidateFunc(_ string) error { return nil }

// NotBlank is a verification function that checks whether the input is empty
func NotBlank(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("input is empty")
	}
	return nil
}

func Done() tea.Msg {
	return DONE
}

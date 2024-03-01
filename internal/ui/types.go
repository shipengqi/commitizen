package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// EchoMode sets the input behavior of the text input field.
// EchoMode is an alias for the textinput.EchoMode.
type EchoMode textinput.EchoMode

// Model is an alias for the tea.Model.
type Model interface {
	tea.Model

	Value() string
	Canceled() bool
}

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
	DefaultTextAreaHeight    = 2
	DefaultSelectWidth       = 20
	DefaultSelectHeight      = 12

	ColorPrompt      = "2"
	colorValidateOk  = "2"
	colorValidateErr = "1"
)

// DefaultValidateFunc is a verification function that does nothing
func DefaultValidateFunc(_ string) error { return nil }

package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	DefaultInputWidth     = 20
	DefaultInputCharLimit = 156
)

type InputModel struct {
	label    string
	canceled bool
	finished bool
	showErr  bool
	init     bool
	err      error

	// validateFunc is a "real-time verification" function, which verifies
	// whether the terminal input data is legal in real time
	validateFunc func(string) error

	// validateOkPrefix is the prompt prefix when the validation fails
	validateOkPrefix string

	// validateErrPrefix is the prompt prefix when the verification is successful
	validateErrPrefix string

	input textinput.Model
}

func NewInput(label string) *InputModel {
	ti := textinput.New()
	ti.CharLimit = DefaultInputCharLimit
	ti.Width = DefaultInputWidth
	ti.EchoMode = textinput.EchoMode(EchoNormal)
	ti.Focus()

	return &InputModel{
		input:             ti,
		label:             label,
		validateFunc:      DefaultValidateFunc,
		validateOkPrefix:  DefaultValidateOkPrefix,
		validateErrPrefix: DefaultValidateErrPrefix,
	}
}

func (m *InputModel) WithPlaceholder(placeholder string) *InputModel {
	m.input.Placeholder = placeholder
	return m
}

func (m *InputModel) WithValidateFunc(fn func(string) error) *InputModel {
	m.validateFunc = fn
	return m
}

func (m *InputModel) WithValidateOkPrefix(prefix string) *InputModel {
	m.validateOkPrefix = prefix
	return m
}

func (m *InputModel) WithValidateErrPrefix(prefix string) *InputModel {
	m.validateErrPrefix = prefix
	return m
}

func (m *InputModel) WithEchoMode(mode EchoMode) *InputModel {
	m.input.EchoMode = textinput.EchoMode(mode)
	return m
}

func (m *InputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *InputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch tmsg := msg.(type) {
	case tea.KeyMsg:
		switch tmsg.Type {
		case tea.KeyCtrlC:
			m.canceled = true
			return m, tea.Quit
		case tea.KeyEnter:
			// If the real-time verification function does not return an error,
			// then the input has been completed
			if m.err == nil {
				m.finished = true
				return m, tea.Quit
			}

			// If there is a verification error, the error message should be display
			m.showErr = true
		case tea.KeyRunes:
			// Hide verification failure message when entering content again
			m.showErr = false
			m.err = nil
		}
		// Call the underlying textinput to update the terminal display
		m.input, cmd = m.input.Update(msg)
		// Perform real-time verification function after each input
		m.err = m.validateFunc(m.input.Value())

	// We handle errors just like any other message
	// Note: msg is error only when there is an unexpected error in the underlying textinput
	case error:
		m.err = tmsg
		m.showErr = true
		return m, nil
	}

	return m, cmd
}

func (m *InputModel) View() string {
	if m.finished {
		switch m.EchoMode() {
		case textinput.EchoNormal:
			return fmt.Sprintf(
				"%s %s\n%s\n",
				FontColor(m.validateOkPrefix, colorValidateOk),
				m.label,
				quitValueStyle.Render(m.Value()),
			)
		case textinput.EchoNone:
			return fmt.Sprintf(
				"%s %s\n",
				FontColor(m.validateOkPrefix, colorValidateOk),
				m.label,
			)
		case textinput.EchoPassword:
			return fmt.Sprintf(
				"%s %s\n%s\n",
				FontColor(m.validateOkPrefix, colorValidateOk),
				m.label,
				quitValueStyle.Render(GenMask(len([]rune(m.Value())))),
			)
		}
	}

	if !m.init {
		m.err = m.validateFunc(m.input.Value())
		m.init = true
	}

	var showMsg, errMsg string
	if m.err != nil {
		showMsg = fmt.Sprintf(
			"%s %s\n%s",
			FontColor(m.validateErrPrefix, colorValidateErr),
			m.label,
			m.input.View(),
		)
		if m.showErr {
			errMsg = FontColor(fmt.Sprintf("%s ERROR: %s\n", m.validateErrPrefix, m.err.Error()), colorValidateErr)
			return fmt.Sprintf("%s\n%s\n", showMsg, errMsg)
		}
	} else {
		showMsg = fmt.Sprintf(
			"%s %s\n%s",
			FontColor(m.validateOkPrefix, colorValidateOk),
			m.label,
			m.input.View(),
		)
	}

	return showMsg + "\n"
}

// Value return the input string
func (m *InputModel) Value() string {
	return m.input.Value()
}

// EchoMode return the input EchoMode
func (m *InputModel) EchoMode() textinput.EchoMode {
	return m.input.EchoMode
}

// Canceled determine whether the operation is cancelled
func (m *InputModel) Canceled() bool {
	return m.canceled
}

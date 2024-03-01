package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var quitValueStyle = lipgloss.NewStyle().Margin(0, 0, 0, 2)

type TextAreaModel struct {
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

	input textarea.Model
}

func NewTextArea(label string) *TextAreaModel {
	ti := textarea.New()
	ti.MaxHeight = DefaultTextAreaMaxHeight
	ti.SetHeight(DefaultTextAreaHeight)
	ti.Focus()

	return &TextAreaModel{
		input:             ti,
		label:             label,
		validateFunc:      DefaultValidateFunc,
		validateOkPrefix:  DefaultValidateOkPrefix,
		validateErrPrefix: DefaultValidateErrPrefix,
	}
}

func (m *TextAreaModel) WithPlaceholder(placeholder string) *TextAreaModel {
	m.input.Placeholder = placeholder
	return m
}

func (m *TextAreaModel) WithWidth(width int) *TextAreaModel {
	m.input.SetWidth(width)
	return m
}

func (m *TextAreaModel) WithHeight(height int) *TextAreaModel {
	if height > m.input.MaxHeight {
		height = m.input.MaxHeight
	}
	m.input.SetHeight(height)
	return m
}

func (m *TextAreaModel) WithMaxHeight(height int) *TextAreaModel {
	m.input.MaxHeight = height
	return m
}

func (m *TextAreaModel) WithValidateFunc(fn func(string) error) *TextAreaModel {
	m.validateFunc = fn
	return m
}

func (m *TextAreaModel) WithValidateOkPrefix(prefix string) *TextAreaModel {
	m.validateOkPrefix = prefix
	return m
}

func (m *TextAreaModel) WithValidateErrPrefix(prefix string) *TextAreaModel {
	m.validateErrPrefix = prefix
	return m
}

func (m *TextAreaModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m *TextAreaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmds []tea.Cmd
	var cmd tea.Cmd

	switch tmsg := msg.(type) {
	case tea.KeyMsg:
		switch tmsg.Type {
		case tea.KeyEsc:
			if m.input.Focused() {
				m.input.Blur()
			}
		case tea.KeyCtrlJ:
			// If the real-time verification function does not return an error,
			// then the input has been completed
			if m.err == nil {
				m.finished = true
				return m, tea.Quit
			}
			// If there is a verification error, the error message should be display
			m.showErr = true
		case tea.KeyCtrlC:
			m.canceled = true
			return m, tea.Quit
		case tea.KeyRunes:
			// Hide verification failure message when entering content again
			m.showErr = false
			m.err = nil
		}
		m.input, cmd = m.input.Update(msg)
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

func (m *TextAreaModel) View() string {
	if m.finished {
		return fmt.Sprintf(
			"%s %s\n%s\n",
			FontColor(m.validateOkPrefix, colorValidateOk),
			m.label,
			quitValueStyle.Render(fmt.Sprintf(m.Value())),
		)
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
func (m *TextAreaModel) Value() string {
	return m.input.Value()
}

// Canceled determine whether the operation is cancelled
func (m *TextAreaModel) Canceled() bool {
	return m.canceled
}

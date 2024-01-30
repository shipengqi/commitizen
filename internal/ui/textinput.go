package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type InputModel struct {
	textInput textinput.Model
	question  string
	err       error
}

func NewInput(question string) *InputModel {
	ti := textinput.New()
	ti.Focus()

	return &InputModel{
		textInput: ti,
		question:  question,
		err:       nil,
	}
}

func (m *InputModel) WithPlaceholder(placeholder string) *InputModel {
	m.textInput.Placeholder = placeholder
	return m
}

func (m *InputModel) WithWidth(width int) *InputModel {
	m.textInput.Width = width
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
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = tmsg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m *InputModel) View() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		m.question,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

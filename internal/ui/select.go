package ui

import (
	"errors"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectModel struct {
	cursor   int
	question string
	choice   string
	choices  []string
}

func NewSelect(question string, choices []string) (SelectModel, error) {
	if question == "" {
		return SelectModel{}, errors.New("")
	}
	if len(choices) == 0 {
		return SelectModel{}, errors.New("")
	}
	return SelectModel{
		question: question,
		choices:  choices,
	}, nil
}

func (m SelectModel) Init() tea.Cmd {
	return nil
}

func (m SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch tmsg := msg.(type) {
	case tea.KeyMsg:
		switch tmsg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		case "enter":
			// Send the choice on the channel and exit.
			m.choice = m.choices[m.cursor]
			return m, tea.Quit
		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}
		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m SelectModel) View() string {
	s := strings.Builder{}
	s.WriteString(m.question)
	s.WriteString("\n\n")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type TextAreaModel struct {
	textarea textarea.Model
	question string
	err      error
}

func NewTextAreaModel(question string) *TextAreaModel {
	ti := textarea.New()
	ti.Focus()

	return &TextAreaModel{
		textarea: ti,
		question: question,
		err:      nil,
	}
}

func (m *TextAreaModel) WithPlaceholder(placeholder string) *TextAreaModel {
	m.textarea.Placeholder = placeholder
	return m
}

func (m *TextAreaModel) WithWidth(width int) *TextAreaModel {
	m.textarea.SetWidth(width)
	return m
}

func (m *TextAreaModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m *TextAreaModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m *TextAreaModel) View() string {
	return fmt.Sprintf(
		"Tell me a story.\n\n%s\n\n%s",
		m.textarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

package ui

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle()
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 2)
)

type Choices []Choice

func (c Choices) toBubblesItem() []list.Item {
	if len(c) == 0 {
		return nil
	}

	var items []list.Item

	for _, v := range c {
		items = append(items, v)
	}
	return items
}

type Choice string

func (i Choice) FilterValue() string { return "" }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Choice)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	_, _ = fmt.Fprint(w, fn(str))
}

type SelectModel struct {
	label    string
	choice   string
	finished bool
	required bool
	showErr  bool
	err      error

	list list.Model
}

func NewSelect(label string, choices Choices) (*SelectModel, error) {
	if label == "" {
		return nil, errors.New("")
	}
	if len(choices) == 0 {
		return nil, errors.New("")
	}
	l := list.New(choices.toBubblesItem(), itemDelegate{}, DefaultSelectWidth, DefaultSelectHeight)
	l.Title = label
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return &SelectModel{list: l, label: label}, nil
}

func (m *SelectModel) WithWidth(width int) *SelectModel {
	m.list.SetWidth(width)
	return m
}

func (m *SelectModel) WithHeight(height int) *SelectModel {
	m.list.SetHeight(height)
	return m
}

func (m *SelectModel) SetRequired(required bool) *SelectModel {
	m.required = required
	return m
}

func (m *SelectModel) Init() tea.Cmd {
	return nil
}

func (m *SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch tmsg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(tmsg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := tmsg.String(); keypress {
		case "q", "ctrl+c":
			if m.required && m.choice == "" {
				m.err = errors.New("this filed is required")
				m.showErr = true
				return m, nil
			}
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(Choice)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		case tea.KeyUp.String(), tea.KeyDown.String(), tea.KeyRight.String(), tea.KeyLeft.String(),
			tea.KeyHome.String(), tea.KeyEnd.String(), tea.KeyPgUp.String(), tea.KeyPgDown.String(),
			"k", "j", "l", "h", "g", "G":
			m.showErr = false
			m.err = nil
		}
		m.list, cmd = m.list.Update(msg)

	case error:
		m.err = tmsg
		m.showErr = true
		return m, nil
	}

	return m, cmd
}

func (m *SelectModel) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s\n%s", m.label, m.choice))
	}
	if m.showErr {
		return fmt.Sprintf("%s\n%s\n", m.list.View(), FontColor(fmt.Sprintf("%s ERROR: %s\n", DefaultValidateErrPrefix, m.err.Error()), colorValidateErr))
	}
	return "\n" + m.list.View()
}

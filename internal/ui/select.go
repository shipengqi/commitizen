package ui

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle()
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingBottom(1)
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
	canceled bool
	err      error

	list list.Model
}

func NewSelect(label string, choices Choices) *SelectModel {
	l := list.New(choices.toBubblesItem(), itemDelegate{}, DefaultSelectWidth, DefaultSelectHeight)
	l.Title = label
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return &SelectModel{list: l, label: label}
}

func (m *SelectModel) WithWidth(width int) *SelectModel {
	m.list.SetWidth(width)
	return m
}

func (m *SelectModel) WithHeight(height int) *SelectModel {
	m.list.SetHeight(height)
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
			m.canceled = true
			return m, tea.Quit
		case "enter":
			i, ok := m.list.SelectedItem().(Choice)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
		m.list, cmd = m.list.Update(msg)

	case error:
		m.err = tmsg
		return m, nil
	}

	return m, cmd
}

func (m *SelectModel) View() string {
	if m.choice != "" {
		// hardcode for the select value
		// fix https://github.com/shipengqi/commitizen/issues/18
		val := m.Value()
		tokens := strings.Split(m.Value(), ":")
		if len(tokens) > 0 {
			val = tokens[0]
		}

		return fmt.Sprintf(
			"%s %s\n%s\n",
			FontColor(DefaultValidateOkPrefix, colorValidateOk),
			m.label,
			quitValueStyle.Render(val),
		)
	}
	return "\n" + m.list.View()
}

func (m *SelectModel) Value() string {
	return m.choice
}

func (m *SelectModel) Canceled() bool {
	return m.canceled
}

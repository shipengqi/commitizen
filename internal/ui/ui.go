package ui

import tea "github.com/charmbracelet/bubbletea"

func Run(model tea.Model, noTTY bool) (tea.Model, error) {
	var p *tea.Program
	if noTTY {
		p = tea.NewProgram(model, tea.WithInput(nil))
	} else {
		p = tea.NewProgram(model)
	}
	return p.Run()
}

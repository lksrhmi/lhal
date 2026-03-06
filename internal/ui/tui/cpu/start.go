package cpu

import tea "github.com/charmbracelet/bubbletea"

func Start() {
	p := tea.NewProgram(New())

	if err := p.Start(); err != nil {
		panic(err)
	}
}
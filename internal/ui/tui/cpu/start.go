package cpu

import tea "github.com/charmbracelet/bubbletea"

func Start() {
	p := tea.NewProgram(New())

	_, runTeaErr := p.Run()
	if runTeaErr != nil {
		panic("TEA got spilled")
	}
}
package cpu

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"cmd/lhal/internal/service/cpu"
	"cmd/lhal/internal/ui/tui/style"
)

type tickMsg float64

type Model struct {
	history []float64
	width   int
	height  int
}

func New() Model {
	return Model{
		history: []float64{},
		width:   60,
		height:  10,
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		usage, _ := cpu.GetUsage()
		return tickMsg(usage)
	})
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tickMsg:
		m.history = append(m.history, float64(msg))

		if len(m.history) > m.width {
			m.history = m.history[1:]
		}

		return m, tick()

	}

	return m, nil
}

func (m Model) View() string {

	graph := drawGraph(m.history, m.height)

	return style.ContainerStyle.Render(
		fmt.Sprintf(
			"%s\n\n%s\n\n%s %.1f%%",
			style.TitleStyle.Render("CPU Monitor"),
			graph,
			style.LabelStyle.Render("Usage:"),
			last(m.history),
		),
	)
}

func last(values []float64) float64 {
	if len(values) == 0 {
		return 0
	}
	return values[len(values)-1]
}
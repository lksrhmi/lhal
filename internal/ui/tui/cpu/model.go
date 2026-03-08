package cpu

import (
	"fmt"
	_ "image/color/palette"
	"time"

	"cmd/lhal/internal/service/cpu"
	"cmd/lhal/internal/ui/tui/style"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg float64

type Model struct {
	history []float64
	width   int
	height  int
}

func New() Model {
	return Model{
		history: []float64{0},
		width:   60,
		height:  10,
	}
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		usage, getUsageErr := cpu.GetUsage()
		if getUsageErr != nil {
			panic("heyho, cpu usage from Service")
		}

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

			// cuts off the last part
			if len(m.history) > m.width {
				m.history = m.history[1:]
			}

			return m, tick()

		case tea.KeyMsg:
			switch msg.String() {
				case "q":
					return m, tea.Quit
			}
	}

	return m, nil
}

func (m Model) View() string {

	graph := getGraph(&m)

	return style.CpuUsageContainerStyle.Render(
		graph,
		fmt.Sprint(
			"\n",
			"q for exit",
		),
	)
}
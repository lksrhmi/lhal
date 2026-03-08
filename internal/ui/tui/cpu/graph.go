package cpu

import (
	"github.com/guptarohit/asciigraph"
)

func getGraph(m *Model) string {
	graph := asciigraph.Plot(
		m.history,
		asciigraph.Height(m.height),
		asciigraph.Width(m.width),
	)

	return graph
}

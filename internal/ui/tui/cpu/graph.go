package cpu

import (
	"strings"
	"cmd/lhal/internal/ui/tui/style"
)

func drawGraph(values []float64, height int) string {

	if len(values) == 0 {
		return ""
	}

	grid := make([][]string, height)

	for i := range grid {
		grid[i] = make([]string, len(values))
		for j := range grid[i] {
			grid[i][j] = " "
		}
	}

	for x, v := range values {

		barHeight := int((v / 100) * float64(height))

		for y := 0; y < barHeight; y++ {

			row := height - 1 - y

			grid[row][x] = style.BarStyle.Render("█")
		}
	}

	var out strings.Builder

	for _, row := range grid {

		for _, cell := range row {
			out.WriteString(cell)
		}

		out.WriteRune('\n')
	}

	return out.String()
}
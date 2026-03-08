package style

import "github.com/charmbracelet/lipgloss"

var CpuUsageContainerStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	Padding(1, 2).
	BorderForeground(lipgloss.Color("63"))

var TitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("86"))

var LabelStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("240"))

var BarStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("212"))
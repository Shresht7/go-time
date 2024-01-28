package laps

import (
	"fmt"

	"github.com/Shresht7/go-time/tui/helpers"
	"github.com/charmbracelet/lipgloss"
)

// VIEW
// ----

func (m Model) View() string {
	s := ""
	for i, lap := range m.laps {
		s += fmt.Sprintf("%s\t%s\t\t%s\n",
			m.styleIndex(i),
			helpers.FormatDuration(lap.duration()),
			helpers.FormatDuration(m.Sum(i)),
		)
	}
	return s
}

// STYLES
var (
	colorRed  = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	colorBlue = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
)

// Color the shortest and longest lap times in red and blue, respectively.
func (m Model) styleIndex(i int) string {
	s := fmt.Sprintf("%d", i+1)

	if i == m.Shortest() {
		return colorBlue.Render(s)
	}
	if i == m.Longest() {
		return colorRed.Render(s)
	}
	return s
}

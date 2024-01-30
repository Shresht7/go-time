package laps

import (
	"fmt"

	"github.com/Shresht7/go-time/pkg/format"
	"github.com/charmbracelet/lipgloss"
)

// STYLES
var (
	header    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("8")).Border(lipgloss.NormalBorder(), false, false, true, false)
	colorRed  = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	colorBlue = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
)

// VIEW
// ----

func (m Model) View() string {
	s := ""

	// If there are no laps, return an empty string
	if len(m.laps) == 0 {
		return s
	}

	// Print the laps table
	s += header.Render("Lap\tLap Time\t\tTotal Time") + "\n"
	for i, lap := range m.laps {
		s += fmt.Sprintf("%s\t%s\t%s\n",
			m.styleIndex(i),
			format.Duration(lap.duration()),
			format.Duration(m.Sum(i)),
		)
	}

	return s
}

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

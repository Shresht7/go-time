package timer

import (
	"github.com/charmbracelet/lipgloss"
)

// VIEW
// ----

func (m Model) View() string {
	var s string
	s += m.FormatTime()
	if m.Running {
		s += " " + m.spinner.View()
	}
	return lipgloss.NewStyle().Padding(1, 2).Render(s)
}

package timer

import (
	"github.com/charmbracelet/lipgloss"
)

// VIEW
// ----

func (m Model) View() string {
	var s string
	if m.Running {
		s += m.spinner.View() + "  "
	} else {
		s += "   "
	}
	s += m.FormatTime()
	return lipgloss.NewStyle().Padding(1, 2).Render(s)
}

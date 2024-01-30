package timer

import (
	"github.com/Shresht7/go-cli-tools/status"
	"github.com/charmbracelet/lipgloss"
)

// VIEW
// ----

func (m Model) View() string {
	var s string
	s += m.FormatTime()
	s += " " + m.ViewProgress()
	return lipgloss.NewStyle().Padding(1, 2).Render(s)
}

func (m Model) ViewProgress() string {
	var s string
	if m.Running {
		s += m.spinner.View()
	} else {
		if m.remaining <= 0 {
			return status.Tick
		} else {
			s += " "
		}
	}
	return s
}

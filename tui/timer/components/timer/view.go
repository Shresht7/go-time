package timer

import (
	"fmt"

	"github.com/Shresht7/go-cli-tools/status"
	"github.com/charmbracelet/lipgloss"
)

// VIEW
// ----

// Returns the timer component view
func (m Model) View() string {
	var s string
	s += m.FormatTime()
	s += " " + m.ViewProgress()
	return lipgloss.NewStyle().Padding(1, 2).Render(s)
}

// Renders the spinner indicating the timer is running
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

// Renders the help menu for the timer
func (m *Model) ViewHelp() string {
	return m.help.View(m.keys)
}

// Formats the remaining time in the format HH:MM:SS
func (m *Model) FormatTime() string {
	hours := m.remaining / 3600
	minutes := (m.remaining - (hours * 3600)) / 60
	seconds := m.remaining - (hours * 3600) - (minutes * 60)
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

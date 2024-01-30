package timer

import (
	"github.com/charmbracelet/lipgloss"
)

// VIEW
// ----

var borderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true)
var focusedStyle = borderStyle.Copy().BorderForeground(lipgloss.Color("12"))
var unfocusedStyle = borderStyle.Copy().BorderForeground(lipgloss.Color("240"))

func (m Model) View() string {
	s := m.FormatTime()
	if m.isFocused {
		return focusedStyle.Render(s)
	}
	return unfocusedStyle.Render(s)
}

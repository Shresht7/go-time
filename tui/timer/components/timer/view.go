package timer

import "github.com/charmbracelet/lipgloss"

// VIEW
// ----

func (m Model) View() string {
	s := m.FormatTime()
	return lipgloss.NewStyle().Padding(1, 2).Render(s)
}

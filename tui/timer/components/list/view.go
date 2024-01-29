package list

import "github.com/charmbracelet/lipgloss"

// VIEW
// ----

var borderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true)
var focusedStyle = borderStyle.Copy().BorderForeground(lipgloss.Color("12"))
var unfocusedStyle = borderStyle.Copy().BorderForeground(lipgloss.Color("240"))

func (m Model) View() string {
	var s string
	if m.prompt != "" {
		s += m.prompt + "\n"
	}
	for i, item := range m.items {
		if i == m.selected {
			s += " ▶ " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	if m.isFocused {
		return focusedStyle.Render(s)
	}
	return unfocusedStyle.Render(s)
}

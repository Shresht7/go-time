package list

import "github.com/charmbracelet/lipgloss"

// VIEW
// ----

var borderStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder(), true).Padding(1, 2)
var focusedStyle = borderStyle.Copy().BorderForeground(lipgloss.Color("12"))
var unfocusedStyle = borderStyle.Copy().BorderForeground(lipgloss.Color("240"))

func (m Model) View() string {

	s := lipgloss.JoinVertical(lipgloss.Top,
		m.ViewPrompt(),
		m.ViewList(),
	)

	if m.isFocused {
		return focusedStyle.Render(s)
	}
	return unfocusedStyle.Render(s)
}

func (m Model) ViewPrompt() string {
	return lipgloss.NewStyle().Bold(true).Render(m.prompt)
}

func (m Model) ViewList() string {
	var s string
	for i, item := range m.items {
		if i == m.selected {
			s += "▶ " + item + "\n"
		} else {
			s += "  " + item + "\n"
		}
	}
	return s
}

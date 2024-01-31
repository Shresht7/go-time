package list

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// STYLES
// ------

var boldStyle = lipgloss.NewStyle().Bold(true)

var selectedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("12"))

var borderStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder(), true).
	BorderForeground(lipgloss.Color("12")).
	Padding(1, 2)

// VIEW
// ----

func (m Model) View() string {
	s := lipgloss.JoinVertical(lipgloss.Top,
		m.ViewPrompt(),
		m.ViewList(),
	)
	return borderStyle.Render(s)
}

func (m Model) ViewPrompt() string {
	return boldStyle.Render(m.prompt)
}

func (m Model) ViewList() string {
	var s string
	for i, item := range m.items {
		if i == m.selected {
			s += "▶ " + selectedStyle.Render(item) + "\n"
		} else {
			s += "  " + item + "\n"
		}
	}
	s = strings.TrimSuffix(s, "\n")
	return s
}

// Returns the help menu for the list component
func (m *Model) ViewHelp() string {
	return m.help.View(m.keys)
}

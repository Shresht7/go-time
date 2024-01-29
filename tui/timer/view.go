package timer

import "github.com/charmbracelet/lipgloss"

// VIEW
// ----

func (m Model) View() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		m.viewComponent(),
		m.viewHelp(),
	)
}

func (m Model) viewComponent() string {
	var s string

	switch m.focused {
	case focusTimer:
		s = m.timer.View()
	case focusList:
		s = m.list.View()
	}

	return s
}

func (m Model) viewHelp() string {
	var help string

	switch m.focused {
	case focusTimer:
		help = m.timer.ViewHelp()
	case focusList:
		help = m.list.ViewHelp()
	}

	help += " • " + m.help.View(m.keys)

	return help
}

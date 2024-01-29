package timer

// VIEW
// ----

func (m Model) View() string {
	s := m.timer.View()
	s += "\n" + m.list.View()
	s += "\n" + m.FocusedHelp()
	return s
}

func (m Model) FocusedHelp() string {
	var s string
	switch m.focused {
	case focusTimer:
		s = m.help.View(m.keys)
	case focusList:
		s = m.list.ViewHelp() + " " + m.help.View(m.keys)
	}
	return s
}

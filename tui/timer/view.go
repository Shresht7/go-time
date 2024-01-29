package timer

// VIEW
// ----

func (m Model) View() string {
	var s string
	switch m.focused {
	case focusTimer:
		s = m.timer.View()
	case focusList:
		s = m.list.View()
	}

	return s
}

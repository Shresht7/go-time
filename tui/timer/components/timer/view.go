package timer

import "strconv"

// VIEW
// ----

func (m Model) View() string {
	s := ""
	s += strconv.Itoa(m.remaining) + "\n"
	s += m.help.View(m.keys)
	return s
}

package timer

import "strconv"

// VIEW
// ----

func (m Model) View() string {
	s := "\n" + strconv.Itoa(m.remaining) + "\n"
	return s
}

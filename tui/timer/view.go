package timer

import "strconv"

// VIEW
// ----

func (m timerModel) View() string {
	return strconv.Itoa(m.remaining)
}

package timer

// VIEW
// ----

func (m timerModel) View() string {
	if m.running {
		return "Timer Running"
	}
	return "Timer Stopped"
}

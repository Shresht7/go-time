package clock

import tea "github.com/charmbracelet/bubbletea"

// ---
// RUN
// ---

func Run() {
	clock := newClockModel()
	p := tea.NewProgram(clock)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

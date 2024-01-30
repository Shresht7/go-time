package clock

import tea "github.com/charmbracelet/bubbletea"

// ---
// RUN
// ---

// Run starts the clock bubble-tea program
func Run() {
	clock := new()
	p := tea.NewProgram(clock)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

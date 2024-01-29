package clock

import tea "github.com/charmbracelet/bubbletea"

// INIT
// ----

// The Init function is called when the program starts. It starts the clock ticker
func (c model) Init() tea.Cmd {
	return tick()
}

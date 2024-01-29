package clock

import tea "github.com/charmbracelet/bubbletea"

// INIT
// ----

func (c clockModel) Init() tea.Cmd {
	return tick()
}

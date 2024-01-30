package timer

import (
	tea "github.com/charmbracelet/bubbletea"
)

// INIT
// ----

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		tick(),
	)
}

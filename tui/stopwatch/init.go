package stopwatch

import tea "github.com/charmbracelet/bubbletea"

// The Init function is called when the program starts and returns a command
// to run immediately. We'll start the stopwatch ticker.
func (m *stopwatchModel) Init() tea.Cmd {
	return tick()
}

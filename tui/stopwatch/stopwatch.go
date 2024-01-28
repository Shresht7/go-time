package stopwatch

import (
	tea "github.com/charmbracelet/bubbletea"
)

// ---
// RUN
// ---

func Run() {
	stopwatch := newStopwatchModel()
	p := tea.NewProgram(stopwatch)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

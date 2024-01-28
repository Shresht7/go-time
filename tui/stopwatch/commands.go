package stopwatch

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// MESSAGES
// --------

// A message to start the stopwatch
type msgStart struct{ t time.Time }

// A message to stop the stopwatch
type msgStop struct{}

type msgLap struct{}

// A message to reset the stopwatch
type msgReset struct{}

// COMMANDS
// --------

// A command to start the stopwatch
func (m *stopwatchModel) cmdStart() tea.Msg {
	return msgStart{time.Now()}
}

// A command to stop the stopwatch
func (m *stopwatchModel) cmdStop() tea.Msg {
	return msgStop{}
}

// A command to reset the stopwatch
func (m *stopwatchModel) cmdReset() tea.Msg {
	return msgReset{}
}

func (m *stopwatchModel) cmdLap() tea.Msg {
	return msgLap{}
}

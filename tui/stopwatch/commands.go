package stopwatch

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// MESSAGES
// --------

// A message to start the stopwatch
type msgStart struct{ t time.Time }

// A message to record a lap time
type msgLap struct{}

// A message to stop the stopwatch
type msgStop struct{}

// A message to reset the stopwatch
type msgReset struct{}

// COMMANDS
// --------

// A command to start the stopwatch
func (m *Model) cmdStart() tea.Msg {
	return msgStart{time.Now()}
}

// A command to record a lap time
func (m *Model) cmdLap() tea.Msg {
	return msgLap{}
}

// A command to stop the stopwatch
func (m *Model) cmdStop() tea.Msg {
	return msgStop{}
}

// A command to reset the stopwatch
func (m *Model) cmdReset() tea.Msg {
	return msgReset{}
}

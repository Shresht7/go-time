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

// A message to resume the stopwatch
type msgResume struct{}

type msgLap struct{}

// A message to reset the stopwatch
type msgReset struct{}

// COMMANDS
// --------

// A command to start the stopwatch
func (m *stopwatchModel) Start() tea.Msg {
	return msgStart{time.Now()}
}

// A command to stop the stopwatch
func (m *stopwatchModel) Stop() tea.Msg {
	return msgStop{}
}

// A command to resume the stopwatch
func (m *stopwatchModel) Resume() tea.Msg {
	return msgResume{}
}

// A command to reset the stopwatch
func (m *stopwatchModel) Reset() tea.Msg {
	return msgReset{}
}

func (m *stopwatchModel) Lap() tea.Msg {
	return msgLap{}
}

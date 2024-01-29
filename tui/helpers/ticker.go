package helpers

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// MESSAGE
// -------

// MsgTick is the message sent by our ticker to update the clock
type MsgTick struct{ T time.Time }

// COMMAND
// -------

// ticker is a helper function that returns a tea.Cmd that sends a MsgTick every interval
func CmdTicker(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(t time.Time) tea.Msg {
		return MsgTick{t}
	})
}

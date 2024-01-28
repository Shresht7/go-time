package stopwatch

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/tui/helpers"
)

// CONSTANTS
// ---------

// The speed at which the stopwatch ticks
const TICK_SPEED = 10 * time.Millisecond

// A tea.Cmd to update the stopwatch every interval
func tick() tea.Cmd {
	return helpers.CmdTicker(TICK_SPEED)
}

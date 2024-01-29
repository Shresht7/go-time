package clock

import (
	"time"

	"github.com/Shresht7/go-time/tui/helpers"
	tea "github.com/charmbracelet/bubbletea"
)

// CONSTANTS
// ---------

// The speed at which the clock ticks
const TICK_SPEED = 1 * time.Second

// A tea.Cmd to update the clock every interval
func tick() tea.Cmd {
	return helpers.CmdTicker(TICK_SPEED)
}

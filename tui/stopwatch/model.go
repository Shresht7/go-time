package stopwatch

import (
	"time"

	"github.com/charmbracelet/bubbles/help"

	"github.com/Shresht7/go-time/tui/stopwatch/components/laps"
)

// MODEL
// -----

// Represents a stopwatch
type Model struct {
	running bool          // Whether the stopwatch is running
	start   time.Time     // The time the stopwatch was started
	end     time.Time     // The time the stopwatch was stopped
	elapsed time.Duration // The elapsed time
	laps    laps.Model    // The lap times

	keys KeyMap     // Key bindings for the stopwatch view
	help help.Model // Help menu model
}

// Creates a new stopwatch model. Initializes the default key bindings and
// help menu. The stopwatch is "stopped" by default.
func New() *Model {
	return &Model{
		keys: DefaultKeyMap,
		help: help.New(),
	}
}

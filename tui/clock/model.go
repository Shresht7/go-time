package clock

import (
	"time"
)

// -----
// MODEL
// -----

// Represents the state of the clock
type model struct {
	t time.Time
}

// Instantiates a clock model
func newClockModel() *model {
	return &model{
		t: time.Now(),
	}
}

// Returns the icon for the clock based on the current time of day
func (c *model) Icon() string {
	hour := c.t.Hour()
	switch {
	case hour >= 6 && hour < 12:
		return "🌄"
	case hour >= 12 && hour < 18:
		return "☀️"
	default:
		return "🌙"
	}
}

package clock

import (
	"fmt"
	"time"
)

// -----
// MODEL
// -----

// Represents the state of the clock
type clockModel struct {
	t time.Time
}

// Returns a new clockModel
func newClockModel() *clockModel {
	return &clockModel{
		t: time.Now(),
	}
}

// Returns the current time in the format "HH:MM:SS"
func (c *clockModel) formatTime() (hours, minutes, seconds string) {
	h := c.t.Hour()
	m := c.t.Minute()
	s := c.t.Second()
	return fmt.Sprintf("%02d", h), fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", s)
}

// Returns the icon for the clock based on the current time of day
func (c *clockModel) Icon() string {
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
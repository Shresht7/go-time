package laps

import "time"

// Represents a single lap
type lap struct {
	start    time.Time
	duration time.Duration
}

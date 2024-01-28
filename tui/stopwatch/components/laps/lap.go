package laps

import "time"

// Represents a single lap
type lap struct {
	start time.Time
	end   time.Time
}

// Calculates the duration of a lap
func (l lap) duration() time.Duration {
	return l.end.Sub(l.start)
}

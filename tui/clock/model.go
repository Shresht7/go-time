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
func new() *model {
	return &model{
		t: time.Now(),
	}
}

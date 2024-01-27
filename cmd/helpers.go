package cmd

import (
	"fmt"
	"time"
)

// Returns the current time in the format "HH:MM:SS"
func getClock(now time.Time) (hours, minutes, seconds string) {
	return fmt.Sprintf("%02d", now.Hour()),
		fmt.Sprintf("%02d", now.Minute()),
		fmt.Sprintf("%02d", now.Second())
}

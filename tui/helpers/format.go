package helpers

import (
	"fmt"
	"time"
)

// Returns the elapsed time as a string in the format "HH:MM:SS:MS"
func FormatElapsed(elapsed time.Duration) string {
	hours := int(elapsed.Hours()) % 24
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60
	millisecond := int(elapsed.Milliseconds() % 1000)
	return fmt.Sprintf("%02dh : %02dm : %02ds : %03dms", hours, minutes, seconds, millisecond)
}

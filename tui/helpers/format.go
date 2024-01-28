package helpers

import (
	"fmt"
	"strings"
	"time"
)

// Returns the elapsed time duration as a string in the format "HH:MM:SS:MS"
func FormatDuration(elapsed time.Duration) string {
	s := []string{}

	hours := int(elapsed.Hours())
	if hours > 0 {
		s = append(s, fmt.Sprintf("%02d", hours))
	}

	minutes := int(elapsed.Minutes()) % 60
	s = append(s, fmt.Sprintf("%02d", minutes))

	seconds := int(elapsed.Seconds()) % 60
	s = append(s, fmt.Sprintf("%02d", seconds))

	microseconds := int(elapsed.Microseconds() % 100)
	s = append(s, fmt.Sprintf("%02d", microseconds))

	return strings.Join(s, ":")
}

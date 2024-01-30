package helpers

import "time"

// Returns a formatted string of the current time and date
func FormatTimeAndDate(t time.Time) string {
	return t.Format("15:04:05  •  Monday, 2 January 2006")
}

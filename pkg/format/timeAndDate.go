package format

import "time"

// Returns a formatted string of the current time and date
func TimeAndDate(t time.Time) string {
	return t.Format("15:04:05  •  Monday, 2 January 2006")
}

// Returns the icon for the clock based on the current time of day
func Icon(t time.Time) string {
	hour := t.Hour()
	switch {
	case hour >= 6 && hour < 12:
		return "🌄"
	case hour >= 12 && hour < 18:
		return "☀️"
	default:
		return "🌙"
	}
}

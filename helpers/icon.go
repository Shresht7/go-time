package helpers

import "time"

// Returns the icon for the clock based on the current time of day
func GetIcon(t time.Time) string {
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

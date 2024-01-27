package cmd

import (
	"fmt"
	"time"
)

// Shows the current time
func ShowTime(now time.Time) {

	hours, minutes, seconds := getClock(now)
	weekday := now.Weekday()
	year, month, day := now.Date()

	fmt.Printf("\n%s:%s:%s  %s  %d %s %d\n\n", hours, minutes, seconds, weekday, day, month, year)

}

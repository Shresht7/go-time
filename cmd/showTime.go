package cmd

import (
	"fmt"
	"time"
)

// Shows the current time
func ShowTime(now time.Time) {

	time := now.Format("15:04:05")
	weekday := now.Weekday()
	year, month, day := now.Date()

	fmt.Printf("\n%s  •  %s  •  %d %s %d\n\n", time, weekday, day, month, year)

}

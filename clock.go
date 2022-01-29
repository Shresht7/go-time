package main

import (
	"fmt"
	"time"
)

//	Show the time
func showClock(now time.Time) {

	//	Get clock and date
	hours, minutes, seconds := getClock(now)
	weekday := now.Weekday()
	year, month, day := now.Date()

	//	Print to stdout
	fmt.Printf("\n%s:%s:%s \t %s \t %d %s %d\n\n", hours, minutes, seconds, weekday, day, month, year)

}

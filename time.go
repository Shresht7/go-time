package main

import (
	"fmt"
	"time"
)

//	Show the time
func showTime(now time.Time) {

	hours, minutes, seconds := getClock(now)
	weekday := now.Weekday()
	year, month, day := now.Date()

	fmt.Printf("\n%s:%s:%s  %s  %d %s %d\n\n", hours, minutes, seconds, weekday, day, month, year)

}

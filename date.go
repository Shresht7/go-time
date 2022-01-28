package main

import (
	"fmt"
	"time"
)

//	Show the date
func showDate(now time.Time) {
	weekday := now.Weekday()
	year, month, day := now.Date()
	fmt.Printf("\n%d %s %d \t %s\n\n", day, month, year, weekday)
}

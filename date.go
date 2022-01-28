package main

import (
	"fmt"
	"time"
)

//	Show the date
func showDate(now time.Time) {
	weekday := now.Weekday()
	year, month, day := now.Date()
	fmt.Printf("\n%s, %d %s %d\n\n", weekday, day, month, year)
}

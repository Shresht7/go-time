package main

import (
	"fmt"
	"time"
)

//	Show the time
func showTime(now time.Time) {

	//	Format hours, minutes, seconds
	h, m, s := now.Clock()
	hour, min, sec := fmt.Sprintf("%d", h), fmt.Sprintf("%d", m), fmt.Sprintf("%d", s)
	if len(hour) == 1 {
		hour = fmt.Sprintf("0%s", hour)
	}
	if len(min) == 1 {
		min = fmt.Sprintf("0%s", min)
	}
	if len(sec) == 1 {
		sec = fmt.Sprintf("0%s", sec)
	}

	//	Format weekday
	weekday := now.Weekday()

	//	Format day, month and year
	year, month, day := now.Date()

	//	Print to stdout
	fmt.Printf("\n%s:%s:%s \t %s \t %d %s %d\n\n", hour, min, sec, weekday, day, month, year)

}

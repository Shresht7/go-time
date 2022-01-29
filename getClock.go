package main

import (
	"fmt"
	"time"
)

func getClock(now time.Time) (hours, minutes, seconds string) {
	h, m, s := now.Clock() //	Get hours, minutes and seconds

	//	Format hours, minutes and seconds
	hours, minutes, seconds = fmt.Sprintf("%d", h), fmt.Sprintf("%d", m), fmt.Sprintf("%d", s)
	if len(hours) == 1 {
		hours = fmt.Sprintf("0%s", hours)
	}
	if len(minutes) == 1 {
		minutes = fmt.Sprintf("0%s", minutes)
	}
	if len(seconds) == 1 {
		seconds = fmt.Sprintf("0%s", seconds)
	}

	return hours, minutes, seconds
}

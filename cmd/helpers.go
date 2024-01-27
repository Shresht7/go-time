package cmd

import (
	"fmt"
	"strconv"
	"time"
)

// Returns the current time in the format "HH:MM:SS"
func getClock(now time.Time) (hours, minutes, seconds string) {
	h, m, s := now.Clock() //	Get hours, minutes and seconds

	//	Format hours, minutes and seconds
	hours, minutes, seconds = strconv.Itoa(h), strconv.Itoa(m), strconv.Itoa(s)
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

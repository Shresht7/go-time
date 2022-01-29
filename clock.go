package main

import (
	"fmt"
	"time"
)

//	Show the clock
func showClock(now time.Time) {
	hours, minutes, seconds := getClock(now)
	fmt.Printf("\n%s:%s:%s\n\n", hours, minutes, seconds)
}

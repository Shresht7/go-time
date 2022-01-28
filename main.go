package main

import (
	"os"
	"time"
)

func main() {
	now := time.Now()

	command := os.Args[1]

	switch command {
	case "time":
		showTime(now)
	case "date":
		showDate(now)
	case "calendar":
		showCalendar(now)
	}
}

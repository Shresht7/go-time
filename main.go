package main

import (
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	command := args[0]

	now := time.Now()

	switch command {
	case "time":
		showTime(now)
	case "date":
		showDate(now)
	case "calendar":
		showCalendar(now)
	}
}

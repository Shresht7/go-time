package main

import (
	"io"
	"os"
	"time"
)

// ---
// RUN
// ---

func run(args []string, stdout io.Writer) error {

	now := time.Now() //	Get the current time

	//	If no command was passed, then show time and exit
	if len(os.Args) < 2 {
		showTime(now)
		os.Exit(0)
	}

	//	Execute the given command
	command := os.Args[1]
	switch command {

	case "clock":
		showClock(now)

	case "date":
		showDate(now)

	case "calendar":
		showCalendar(now)

	case "now":
		showTime(now)

	case "time":
		showTime(now)

	default:
		showTime(now)

	}

	// 	Exit with success
	return nil
}

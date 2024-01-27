package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// ----
// MAIN
// ----

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// ---
// RUN
// ---

func run(args []string, stdout io.Writer) error {

	now := time.Now() //	Get the current time

	//	If no command was passed, then show time and exit.
	if len(os.Args) < 2 {
		showTime(now)
		os.Exit(0)
	}

	//	Execute the given command
	command := os.Args[1]
	switch command {

	//	-----
	//	CLOCK
	//	-----

	case "clock":
		showClock(now)

	//	----
	//	DATE
	//	----

	case "date":
		showDate(now)

	//	--------
	//	CALENDAR
	//	--------

	case "calendar":
		showCalendar(now)

	//	----
	//	TIME
	//	----

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

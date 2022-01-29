package main

import (
	"os"
	"time"
)

func main() {
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

}

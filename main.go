package main

import (
	"os"
	"time"
)

func main() {
	now := time.Now() //	Get the current time

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

package main

import (
	"io"
	"os"
	"time"

	"github.com/Shresht7/go-time/cmd"
	"github.com/Shresht7/go-time/tui"
)

// ---
// RUN
// ---

func run(args []string, stdout io.Writer) error {

	now := time.Now() //	Get the current time

	//	If no command was passed, then show time and exit
	if len(os.Args) < 2 {
		cmd.ShowTime(now)
		os.Exit(0)
	}

	//	Execute the given command
	command := os.Args[1]
	switch command {

	case "now":
		cmd.ShowTime(now)

	case "time":
		cmd.ShowTime(now)

	case "clock":
		tui.ShowClock()

	case "date":
		cmd.ShowDate(now)

	case "calendar":
		cmd.ShowCalendar(now)

	default:
		cmd.ShowTime(now)

	}

	// 	Exit with success
	return nil
}

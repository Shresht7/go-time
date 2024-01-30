package main

import (
	"io"
	"os"
	"time"

	"github.com/Shresht7/go-time/cmd"
	"github.com/Shresht7/go-time/tui/clock"
	"github.com/Shresht7/go-time/tui/stopwatch"
	"github.com/Shresht7/go-time/tui/timer"
)

// ---
// RUN
// ---

func run(args []string, stdout io.Writer) error {

	now := time.Now() //	Get the current time

	//	If no command was passed, then show time and exit
	if len(os.Args) < 2 {
		cmd.ShowTime()
		os.Exit(0)
	}

	//	Execute the given command
	command := os.Args[1]
	switch command {

	case "now":
		cmd.ShowTime()

	case "time":
		cmd.ShowTime()

	case "clock":
		clock.Run()

	case "date":
		cmd.ShowDate(now)

	case "calendar":
		cmd.ShowCalendar(now)

	case "stopwatch":
		stopwatch.Run()

	case "timer":
		timer.Run()

	default:
		cmd.ShowTime()

	}

	// 	Exit with success
	return nil
}

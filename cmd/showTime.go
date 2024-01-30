package cmd

import (
	"fmt"
	"time"

	"github.com/Shresht7/go-time/helpers"
)

// Shows the current time
func ShowTime(now time.Time) {
	icon := helpers.GetIcon(now)
	timeAndDate := helpers.FormatTimeAndDate(now)
	fmt.Printf("\n%s   %s\n\n", icon, timeAndDate)
}

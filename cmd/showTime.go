package cmd

import (
	"fmt"
	"time"

	"github.com/Shresht7/go-time/pkg/format"
)

// Shows the current time and date
func ShowTime(now time.Time) {
	icon := format.Icon(now)
	timeAndDate := format.TimeAndDate(now)
	result := icon + "   " + timeAndDate
	fmt.Println("\n" + result + "\n")
}

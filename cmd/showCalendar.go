package cmd

import (
	"fmt"
	"time"

	"github.com/Shresht7/go-time/pkg/calendar"
)

// Show calendar command
func ShowCalendar() {
	now := time.Now()
	calendar := calendar.Render(now)
	fmt.Print(calendar)
}

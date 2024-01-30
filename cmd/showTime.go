package cmd

import (
	"fmt"
	"time"

	"github.com/Shresht7/go-time/pkg/format"
)

// Shows the current time and date
func ShowTime() {
	now := time.Now()
	result := fmt.Sprintf("%s   %s", format.Icon(now), format.TimeAndDate(now))
	fmt.Println("\n" + result + "\n")
}

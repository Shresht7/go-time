package cmd

import (
	"fmt"
	"time"
)

// Shows the current date
func ShowDate() {
	now := time.Now()
	date := now.Format("Monday, 2 January 2006")
	fmt.Println("\n" + date + "\n")
}

package cmd

import (
	"fmt"
	"time"

	"github.com/Shresht7/go-time/pkg/format"
	"github.com/Shresht7/go-time/pkg/layout"
)

// Shows the current time and date
func ShowTime() {
	now := time.Now()
	result := layout.Column(
		layout.Row(format.Icon(now), format.TimeAndDate(now)).WithGap(3).String(),
	)
	fmt.Println(result)
}

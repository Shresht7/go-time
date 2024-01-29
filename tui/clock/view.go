package clock

import (
	"github.com/Shresht7/go-time/tui/helpers"
)

// VIEW
// ----

// The View function is called every time the model is updated
func (c *model) View() string {
	s := helpers.FlexBoxRow(
		c.Icon(),                             // add the icon for the current time of day
		c.t.Format("15:04:05"),               // add the time in the format "HH:MM:SS"
		c.t.Format("Monday, 2 January 2006"), // add the date in the format "Saturday, 27 January 2024"
	).WithSeparator("\t").WithPadding(1).String()
	return s
}

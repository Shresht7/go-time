package clock

import (
	"github.com/Shresht7/go-time/helpers"
)

// VIEW
// ----

// The View function is called every time the model is updated
func (c *model) View() string {
	icon := helpers.GetIcon(c.t)
	timeAndDate := helpers.FormatTimeAndDate(c.t)
	return "\n" + icon + "   " + timeAndDate + "\n\n"
}

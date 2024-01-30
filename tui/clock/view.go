package clock

import (
	"github.com/Shresht7/go-time/pkg/format"
)

// VIEW
// ----

// The View function is called every time the model is updated
func (c *model) View() string {
	icon := format.Icon(c.t)
	timeAndDate := format.TimeAndDate(c.t)
	result := icon + "   " + timeAndDate
	return "\n" + result + "\n\n"
}

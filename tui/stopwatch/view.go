package stopwatch

import (
	"github.com/Shresht7/go-time/tui/helpers"
)

// VIEW
// ----

// The View function is responsible for rendering the UI.
func (m *stopwatchModel) View() string {
	s := helpers.FlexBoxColumn(helpers.Filter(
		"\n"+helpers.FormatDuration(m.elapsed)+"\n",
		m.laps.View(),
		m.help.View(m.keys),
	)...).String()

	return s
}

package stopwatch

import (
	"github.com/Shresht7/go-time/pkg/format"
	"github.com/Shresht7/go-time/pkg/layout"
	"github.com/Shresht7/go-time/tui/helpers"
)

// VIEW
// ----

// The View function is responsible for rendering the UI.
func (m *stopwatchModel) View() string {
	s := layout.Column(helpers.Filter(
		"\n"+format.Duration(m.elapsed)+"\n",
		m.laps.View(),
		m.help.View(m.keys),
	)...).String()

	return s
}

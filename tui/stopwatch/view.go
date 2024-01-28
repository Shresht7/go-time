package stopwatch

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"

	"github.com/Shresht7/go-time/tui/helpers"
)

// VIEW
// ----

var styles = lipgloss.NewStyle().Padding(1, 0)

// The View function is responsible for rendering the UI.
func (m *stopwatchModel) View() string {

	laps := ""
	if len(m.laps) > 0 {
		laps = "Laps\tLap Time\n"
		for i, lap := range m.laps {
			laps += fmt.Sprintf("%d\t%s\n", i+1, helpers.FormatDuration(lap))
		}
	}

	s := lipgloss.JoinVertical(lipgloss.Top, helpers.Filter(
		styles.Render(helpers.FormatDuration(m.elapsed)),
		laps,
		m.help.View(m.keys),
	)...)

	return s
}

package stopwatch

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/Shresht7/go-time/tui/helpers"
)

// VIEW
// ----

var styles = lipgloss.NewStyle().Padding(1, 0)

// The View function is responsible for rendering the UI.
func (m *stopwatchModel) View() string {
	s := lipgloss.JoinVertical(lipgloss.Top, helpers.Filter(
		styles.Render(helpers.FormatDuration(m.elapsed)),
		m.laps.View(),
		m.help.View(m.keys),
	)...)

	return s
}

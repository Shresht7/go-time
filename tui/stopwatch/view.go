package stopwatch

import (
	"github.com/charmbracelet/lipgloss"

	"github.com/Shresht7/go-time/pkg/format"
	"github.com/Shresht7/go-time/tui/helpers"
)

// VIEW
// ----

// The View function is responsible for rendering the UI.
func (m *Model) View() string {
	s := lipgloss.JoinVertical(lipgloss.Left, helpers.Filter(
		"\n"+format.Duration(m.elapsed)+"\n",
		m.laps.View(),
		m.help.View(m.keys),
	)...)
	return s
}

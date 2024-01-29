package timer

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/tui/timer/components/timer"
)

// RUN
// ---

func Run() {
	timer := timer.New()
	p := tea.NewProgram(timer)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

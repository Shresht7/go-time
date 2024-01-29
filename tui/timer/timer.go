package timer

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/tui/timer/components/timer"
)

// RUN
// ---

func Run() {
	timer := timer.NewTimerModel()
	p := tea.NewProgram(timer)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

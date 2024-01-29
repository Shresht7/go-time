package timer

import (
	tea "github.com/charmbracelet/bubbletea"
)

// RUN
// ---

func Run() {
	timer := New()
	p := tea.NewProgram(timer)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

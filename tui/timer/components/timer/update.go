package timer

import (
	"github.com/Shresht7/go-time/tui/helpers"
	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
// ------

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg.(type) {
	case helpers.MsgTick:
		if m.Running {
			m.remaining = m.remaining - 1
		}
		return m, tick()
	}
	return m, nil
}

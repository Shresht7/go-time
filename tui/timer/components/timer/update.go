package timer

import (
	"github.com/Shresht7/go-time/tui/helpers"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
// ------

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case helpers.MsgTick:
		if m.Running {
			m.remaining = m.remaining - 1
		}
		return m, tick()
	case tea.KeyMsg:
		switch {
		// Keypress: Spacebar
		case key.Matches(msg, m.keys.Space):
			if !m.Running {
				m.Start()
				m.keys.Space.SetHelp("spacebar", "pause")
			} else {
				m.Stop()
				m.keys.Space.SetHelp("spacebar", "resume")
			}
		}
	}

	return m, nil
}

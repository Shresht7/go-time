package timer

import (
	"github.com/Shresht7/go-time/tui/helpers"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
// ------

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window Resize Event
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	case helpers.MsgTick:
		if m.running {
			m.remaining = m.remaining - 1
		}
		return m, tick()

	case tea.KeyMsg:
		switch {

		// Switch
		case key.Matches(msg, m.keys.Space):
			if m.running {
				m.keys.Space.SetHelp("<spacebar>", "start")
			} else {
				m.keys.Space.SetHelp("<spacebar>", "stop")
			}
			m.running = !m.running

		// Quit
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		}

	}

	return m, nil
}

package calendar

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window Resize Event
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	case tea.KeyMsg:
		switch {

		// Previous month
		case key.Matches(msg, m.keys.Left):
			m.t = m.t.AddDate(0, -1, 0)

		// Next month
		case key.Matches(msg, m.keys.Right):
			m.t = m.t.AddDate(0, 1, 0)

		// Quit
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		}

	}

	return m, nil
}

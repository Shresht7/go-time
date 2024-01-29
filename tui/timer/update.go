package timer

import (
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

	// Key Press Event
	case tea.KeyMsg:
		switch {

		// Keypress: Tab
		case key.Matches(msg, m.keys.Tab):
			switch m.focused {
			case focusTimer:
				m.SetFocus(focusList)
			case focusList:
				m.SetFocus(focusTimer)
			}

		// Keypress: Spacebar
		case key.Matches(msg, m.keys.Space):
			if !m.timer.Running {
				m.timer.Start()
				m.keys.Space.SetHelp("<spacebar>", "pause")
			} else {
				m.timer.Stop()
				m.keys.Space.SetHelp("<spacebar>", "resume")
			}

		// Keypress: Quit
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

	}

	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch m.focused {
	case focusTimer:
		m.timer, cmd = m.timer.Update(msg)
		cmds = append(cmds, cmd)
	case focusList:
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

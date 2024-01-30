package timer

import (
	"github.com/Shresht7/go-time/tui/helpers"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
// ------

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case helpers.MsgTick:
		if m.Running {
			m.remaining = m.remaining - 1
			if m.remaining <= 0 {
				m.Stop()
			}
		}
		cmds = append(cmds, tick())

	case tea.KeyMsg:

		if !m.isFocused {
			return m, nil
		}

		switch {

		// Keypress: Spacebar
		case key.Matches(msg, m.keys.Space):
			if !m.Running {
				m.Start()
			} else {
				m.Stop()
			}

		// Keypress: r
		case key.Matches(msg, m.keys.Reset):
			m.Reset()
			m.Stop()

		}
	}

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

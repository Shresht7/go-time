package timer

import (
	"github.com/Shresht7/go-time/tui/timer/components/list"
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

	// Preset Selection Event
	case list.MsgSelect:
		preset := m.list.At(msg.Index)
		m.timer.Set(preset)
		m.SetFocus(focusTimer)

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

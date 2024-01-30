package list

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
// ------

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Key Press
	case tea.KeyMsg:

		if !m.isFocused {
			return m, nil
		}

		switch {

		// Select
		case key.Matches(msg, m.keys.Select):
			return m, m.CmdSelect()

		// Previous
		case key.Matches(msg, m.keys.Previous):
			m.Previous()

		// Next
		case key.Matches(msg, m.keys.Next):
			m.Next()

		}

	}

	return m, nil
}

// A message to select the currently selected item
type MsgSelect struct{ Index int }

// A command to select the currently selected item
func (m *Model) CmdSelect() tea.Cmd {
	return func() tea.Msg {
		return MsgSelect{m.selected}
	}
}

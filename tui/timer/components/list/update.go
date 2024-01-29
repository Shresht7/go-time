package list

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
// ------

// Updates the list component.
func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Key Press
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Previous):
			m.Previous()
		case key.Matches(msg, m.keys.Next):
			m.Next()
		}
	}

	return m, nil
}

package timer

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/tui/timer/components/list"
)

// UPDATE
// ------

// Handles events and messages for the timer application
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window Resize Event
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate its view as needed
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

	var cmd tea.Cmd    // Temporary variable for holding commands returned by subcomponents' Update methods
	var cmds []tea.Cmd // The collection of commands to return to the Bubble Tea runtime

	// Update the subcomponent's model and collect its command
	m.timer, cmd = m.timer.Update(msg)
	cmds = append(cmds, cmd)
	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	// Return the updated model along with the batched commands
	return m, tea.Batch(cmds...)
}

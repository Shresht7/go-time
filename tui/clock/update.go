package clock

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/tui/helpers"
)

// UPDATE
// ------

// The Update function is called when events happen, like keypresses
func (c *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Message: clock tick - update the time
	case helpers.MsgTick:
		c.t = msg.T
		return c, tick()

	// Message: keypress
	case tea.KeyMsg:
		switch msg.String() {

		// Keypress: quit
		case "q", "esc", "ctrl+c":
			return c, tea.Quit

		}

	}

	return c, nil
}

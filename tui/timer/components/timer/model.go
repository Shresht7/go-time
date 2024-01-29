package timer

import "github.com/charmbracelet/bubbles/help"

// MODEL
// -----

type Model struct {
	Running   bool // Whether the timer is running
	remaining int  // Remaining time in seconds

	isFocused bool // Whether the timer is focused

	keys KeyMap     // Keybindings model
	help help.Model // Help model
}

func New() Model {
	return Model{
		remaining: 72,
		keys:      DefaultKeyMap,
		help:      help.New(),
	}
}

func (m *Model) Start() {
	m.Running = true
}

func (m *Model) Stop() {
	m.Running = false
}

// Set the focus state of the timer
func (m *Model) SetFocused(f bool) {
	m.isFocused = f
}

func (m *Model) ViewHelp() string {
	return m.help.View(m.keys)
}

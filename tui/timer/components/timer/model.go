package timer

import "github.com/charmbracelet/bubbles/help"

// MODEL
// -----

type Model struct {
	running   bool // Whether the timer is running
	remaining int  // Remaining time in seconds

	keys KeyMap     // Key bindings for the timer view
	help help.Model // Help menu model
}

func New() Model {
	return Model{
		remaining: 72,
		keys:      DefaultKeyMap,
		help:      help.New(),
	}
}

func (m *Model) Start() {
	m.running = true
}

func (m *Model) Stop() {
	m.running = false
}

package timer

import "github.com/charmbracelet/bubbles/help"

// MODEL
// -----

type timerModel struct {
	running   bool // Whether the timer is running
	remaining int  // Remaining time in seconds

	keys KeyMap     // Key bindings for the timer view
	help help.Model // Help menu model
}

func newTimerModel() timerModel {
	return timerModel{
		remaining: 72,
		keys:      DefaultKeyMap,
		help:      help.New(),
	}
}

func (m *timerModel) Start() {
	m.running = true
}

func (m *timerModel) Stop() {
	m.running = false
}

package timer

// MODEL
// -----

type Model struct {
	Running   bool // Whether the timer is running
	remaining int  // Remaining time in seconds

	isFocused bool // Whether the timer is focused
}

func New() Model {
	return Model{
		remaining: 72,
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

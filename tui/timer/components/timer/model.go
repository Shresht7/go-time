package timer

// MODEL
// -----

type Model struct {
	Running   bool // Whether the timer is running
	remaining int  // Remaining time in seconds
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

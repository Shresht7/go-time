package timer

// MODEL
// -----

type timerModel struct {
	running   bool // Whether the timer is running
	remaining int  // Remaining time in seconds
}

func newTimerModel() timerModel {
	return timerModel{
		remaining: 72,
	}
}

func (m *timerModel) Start() {
	m.running = true
}

func (m *timerModel) Stop() {
	m.running = false
}

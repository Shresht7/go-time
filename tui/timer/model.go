package timer

// MODEL
// -----

type timerModel struct {
	running bool // Whether the timer is running
}

func newTimerModel() timerModel {
	return timerModel{}
}

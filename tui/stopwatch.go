package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// CONSTANTS
// ---------

// The speed at which the clock ticks
const TICK_SPEED = 10 * time.Millisecond

// A tea.Cmd to update the clock every interval
func tick() tea.Cmd {
	return ticker(TICK_SPEED)
}

// MODEL
// -----

type stopwatchModel struct {
	running bool
	start   time.Time
	end     time.Time
	elapsed time.Duration
}

// INIT, UPDATE, VIEW
// ------------------

func (m *stopwatchModel) Init() tea.Cmd {
	return tick()
}

func (m *stopwatchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case msgStart:
		m.start = msg.t
		m.running = true
		return m, nil

	case msgStop:
		m.end = time.Now()
		m.running = false
		return m, nil

	case msgTick:
		if m.running {
			m.elapsed = time.Since(m.start)
		}
		return m, tick()

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case " ":
			if m.running {
				return m, m.Stop
			} else {
				return m, m.Start
			}
		}
	}

	return m, nil
}

func (m *stopwatchModel) View() string {
	s := "\n"

	s += m.elapsed.Truncate(TICK_SPEED).String()

	s += "\n\n"

	if m.running {
		s += "Press <space> to stop"
	} else {
		s += "Press <space> to start"
	}

	s += "| Press <q> to quit"

	return s
}

// -------
// COMMAND
// -------

func ShowStopwatch() {
	p := tea.NewProgram(&stopwatchModel{})
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

// MESSAGES & COMMANDS
// -------------------

// A message to start the stopwatch
type msgStart struct{ t time.Time }

// A message to stop the stopwatch
type msgStop struct{}

// A command to start the stopwatch
func (m *stopwatchModel) Start() tea.Msg {
	return msgStart{time.Now()}
}

// A command to stop the stopwatch
func (m *stopwatchModel) Stop() tea.Msg {
	return msgStop{}
}

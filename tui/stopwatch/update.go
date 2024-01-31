package stopwatch

import (
	"time"

	"github.com/Shresht7/go-time/tui/helpers"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// UPDATE
// ------

// The Update function is called when events happen, like keypresses
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window Resize Event
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	// Stopwatch Start
	case msgStart:
		m.start = time.Now().Add(-m.elapsed)
		m.running = true
		m.keys.Space.SetHelp("spacebar", "stop")
		return m, nil

	// Stopwatch Stop
	case msgStop:
		m.end = time.Now()
		m.running = false
		m.keys.Space.SetHelp("spacebar", "start")
		return m, nil

	// Stopwatch Lap
	case msgLap:
		start := m.laps.LastOr(m.start)
		duration := time.Since(start)
		m.laps.Add(start, duration)
		return m, nil

	// Stopwatch Reset
	case msgReset:
		m.elapsed = 0
		m.running = false
		m.laps.Reset()
		m.keys.Space.SetHelp("spacebar", "start")
		return m, nil

	// Stopwatch Tick
	case helpers.MsgTick:
		if m.running {
			m.elapsed = time.Since(m.start)
		}
		return m, tick()

	// Key Press
	case tea.KeyMsg:
		switch {

		// Quit
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit

		// Space
		case key.Matches(msg, DefaultKeyMap.Space):
			if m.running {
				return m, m.cmdStop
			} else {
				return m, m.cmdStart
			}

		// R
		case key.Matches(msg, DefaultKeyMap.R):
			return m, m.cmdReset

		// Enter
		case key.Matches(msg, DefaultKeyMap.Enter):
			return m, m.cmdLap

		}

	}

	return m, nil
}

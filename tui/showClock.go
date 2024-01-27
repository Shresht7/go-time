package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// CLOCK MODEL
// -----------

// Represents a clock
type ClockModel struct {
	t time.Time
}

// Returns the current time in the format "HH:MM:SS"
func (c *ClockModel) formatTime() (hours, minutes, seconds string) {
	h := c.t.Hour()
	m := c.t.Minute()
	s := c.t.Second()
	return fmt.Sprintf("%02d", h), fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", s)
}

// Returns the icon for the clock based on the current time of day
func (c *ClockModel) Icon() string {
	hour := c.t.Hour()
	switch {
	case hour >= 6 && hour < 12:
		return "🌄"
	case hour >= 12 && hour < 18:
		return "☀️"
	default:
		return "🌙"
	}
}

// INIT, UPDATE, VIEW
// ------------------

// The Init function is called when the program starts
func (c *ClockModel) Init() tea.Cmd {
	c.t = time.Now()           // initialize the time immediately on startup
	return ticker(time.Second) // start the ticker to update the clock every second
}

// The Update function is called when events happen, like keypresses
func (c *ClockModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgTick:
		c.t = msg.time
		return c, ticker(time.Second)
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return c, tea.Quit
		}
	}
	return c, nil
}

// The View function is called every time the model is updated
func (c *ClockModel) View() string {
	s := "\n" // start with a newline for padding

	s += c.Icon() // add the icon for the current time of day

	// add the time in the format "HH:MM:SS"
	hours, minutes, seconds := c.formatTime()
	s += "  " + hours + ":" + minutes + ":" + seconds + "\t"

	// add the date in the format "Saturday, 27 January 2024"
	s += c.t.Format("Monday, 2 January 2006")

	// add a newline and some padding at the bottom
	s += "\n\n"

	return s
}

// SHOW CLOCK COMMAND
// ------------------

func ShowClock() {
	p := tea.NewProgram(&ClockModel{})
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

// MESSAGE & COMMAND
// -----------------

// msgTick is the message sent by our ticker to update the clock.
type msgTick struct{ time time.Time }

// ticker is a helper function that returns a tea.Cmd that sends a msgTick every interval.
func ticker(interval time.Duration) tea.Cmd {
	return tea.Tick(interval, func(t time.Time) tea.Msg {
		return msgTick{time: t}
	})
}

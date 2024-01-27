package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// BUBBLE TEA MODEL & METHODS
// --------------------------

type ClockModel struct {
	t time.Time
}

func (c *ClockModel) formatTime() (hours, minutes, seconds string) {
	h := c.t.Hour()
	m := c.t.Minute()
	s := c.t.Second()
	return fmt.Sprintf("%02d", h), fmt.Sprintf("%02d", m), fmt.Sprintf("%02d", s)
}

func (c *ClockModel) Init() tea.Cmd {
	c.t = time.Now()           // initialize the time immediately on startup
	return ticker(time.Second) // start the ticker to update the clock every second
}

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

func (c *ClockModel) View() string {
	hours, minutes, seconds := c.formatTime()
	return "\n" + hours + ":" + minutes + ":" + seconds + "\n\n"
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

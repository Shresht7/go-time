package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// BUBBLE TEA MODEL & METHODS
// --------------------------

type ClockModel struct {
	Hours   string
	Minutes string
	Seconds string
}

func (c *ClockModel) Init() tea.Cmd {
	return ticker(time.Second)
}

func (c *ClockModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case msgTick:
		c.Hours = fmt.Sprintf("%2d", msg.time.Hour())
		c.Minutes = fmt.Sprintf("%2d", msg.time.Minute())
		c.Seconds = fmt.Sprintf("%2d", msg.time.Second())
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
	return "\n" + c.Hours + ":" + c.Minutes + ":" + c.Seconds + "\n\n"
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

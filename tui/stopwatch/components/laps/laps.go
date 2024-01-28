package laps

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/tui/helpers"
)

// MODEL
// -----

// Represents a list of lap times
type Model struct {
	laps []time.Duration
}

// Creates a new lap times model
func New() Model {
	return Model{}
}

// Adds a lap time to the list
func (m *Model) Add(d time.Duration) {
	m.laps = append(m.laps, d)
}

// Clears the list of lap times
func (m *Model) Reset() {
	m.laps = []time.Duration{}
}

// Returns the number of lap times in the list
func (m Model) Len() int {
	return len(m.laps)
}

// INIT
// ----

func (m Model) Init() tea.Cmd {
	return nil
}

// UPDATE
// ------

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

// VIEW
// ----

func (m Model) View() string {
	s := ""
	for i, lap := range m.laps {
		s += fmt.Sprintf("%d\t%s\n", i+1, helpers.FormatDuration(lap))
	}
	return s
}

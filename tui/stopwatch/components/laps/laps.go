package laps

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/Shresht7/go-time/tui/helpers"
)

// MODEL
// -----

// Represents a list of lap times
type Model struct {
	laps []time.Duration // The collection of lap times
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

func (m Model) Shortest() int {
	shortest := 0
	for i, lap := range m.laps {
		if lap < m.laps[shortest] {
			shortest = i
		}
	}
	return shortest
}

func (m Model) Longest() int {
	longest := 0
	for i, lap := range m.laps {
		if lap > m.laps[longest] {
			longest = i
		}
	}
	return longest
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
		s += fmt.Sprintf("%s\t%s\n", m.styleIndex(i), helpers.FormatDuration(lap))
	}
	return s
}

var colorRed = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
var colorBlue = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))

func (m Model) styleIndex(i int) string {
	s := fmt.Sprintf("%d", i+1)

	if i == m.Shortest() {
		return colorRed.Render(s)
	}
	if i == m.Longest() {
		return colorBlue.Render(s)
	}
	return s
}

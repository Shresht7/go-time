package calendar

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
)

type Model struct {
	t time.Time

	keys KeyMaps
	help help.Model
}

func New() Model {
	return Model{
		t:    time.Now(),
		keys: DefaultKeyMaps,
		help: help.New(),
	}
}

// Get the first day of the month
func (m Model) FirstDay() time.Time {
	return time.Date(m.t.Year(), m.t.Month(), 1, 0, 0, 0, 0, time.UTC)
}

// Go to the next month
func (m *Model) NextMonth() {
	t := m.FirstDay()
	m.t = t.AddDate(0, 1, 0) // Add 1 month
}

// Go to the previous month
func (m *Model) PrevMonth() {
	t := m.FirstDay()
	m.t = t.AddDate(0, -1, 0) // Subtract 1 month
}

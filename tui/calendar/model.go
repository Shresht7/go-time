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

func (m *Model) NextMonth() {
	t := time.Date(m.t.Year(), m.t.Month(), 1, 0, 0, 0, 0, time.UTC)
	m.t = t.AddDate(0, 1, 0)
}

func (m *Model) PrevMonth() {
	t := time.Date(m.t.Year(), m.t.Month(), 1, 0, 0, 0, 0, time.UTC)
	m.t = t.AddDate(0, -1, 0)
}

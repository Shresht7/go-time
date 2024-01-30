package calendar

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/pkg/calendar"
)

type Model struct {
	t time.Time
}

func New() Model {
	return Model{
		t: time.Now(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "left", "h":
			m.t = m.t.AddDate(0, -1, 0)

		case "right", "l":
			m.t = m.t.AddDate(0, 1, 0)

		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		}

	}

	return m, nil
}

func (m Model) View() string {
	return calendar.Render(m.t)
}

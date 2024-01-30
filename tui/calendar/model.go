package calendar

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/pkg/calendar"
)

type Model struct {
	cal *calendar.Month
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	return m.cal.String()
}

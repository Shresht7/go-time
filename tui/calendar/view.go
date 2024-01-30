package calendar

import "github.com/Shresht7/go-time/pkg/calendar"

func (m Model) View() string {
	return calendar.Render(m.t)
}

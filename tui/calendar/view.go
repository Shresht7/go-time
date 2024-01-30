package calendar

import "github.com/Shresht7/go-time/pkg/calendar"

func (m Model) View() string {
	s := ""
	s += calendar.Render(m.t)
	s += m.help.View(m.keys)
	return s
}

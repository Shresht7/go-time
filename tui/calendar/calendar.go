package calendar

import tea "github.com/charmbracelet/bubbletea"

func Run() {
	calender := New()
	p := tea.NewProgram(calender)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

package timer

import (
	tea "github.com/charmbracelet/bubbletea"
)

// RUN
// ---

func Run() {
	timer := newTimerModel()
	p := tea.NewProgram(timer)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

type timerModel struct {
	running bool // Whether the timer is running
}

func newTimerModel() timerModel {
	return timerModel{}
}

func (m timerModel) Init() tea.Cmd {
	return nil
}

func (m timerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		// Switch
		case " ":
			m.running = !m.running

		// Quit
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		}

	}

	return m, nil
}

func (m timerModel) View() string {
	if m.running {
		return "Timer Running"
	}
	return "Timer Stopped"
}

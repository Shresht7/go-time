package calendar

import tea "github.com/charmbracelet/bubbletea"

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

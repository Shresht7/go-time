package timer

import tea "github.com/charmbracelet/bubbletea"

// UPDATE
// ------

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

package timer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Shresht7/go-time/tui/timer/components/timer"
)

// RUN
// ---

func Run() {
	timer := New()
	p := tea.NewProgram(timer)
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

// MODEL
// -----

type Model struct {
	timer timer.Model // The timer component

	keys KeyMap     // Key bindings model
	help help.Model // Help menu model
}

func New() Model {
	return Model{
		timer: timer.New(),
		keys:  DefaultKeyMap,
		help:  help.New(),
	}
}

// INIT
// ----

func (m Model) Init() tea.Cmd {
	return m.timer.Init()
}

// UPDATE
// ------

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	// Window Resize Event
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	// Key Press Event
	case tea.KeyMsg:
		switch {

		// Keypress: Spacebar
		case key.Matches(msg, m.keys.Space):
			if !m.timer.Running {
				m.timer.Start()
				m.keys.Space.SetHelp("<spacebar>", "pause")
			} else {
				m.timer.Stop()
				m.keys.Space.SetHelp("<spacebar>", "resume")
			}

		// Keypress: Quit
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}

	}

	var cmd tea.Cmd
	var cmds []tea.Cmd
	m.timer, cmd = m.timer.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

// VIEW
// ----

func (m Model) View() string {
	s := m.timer.View()
	s += "\n" + m.help.View(m.keys)
	return s
}

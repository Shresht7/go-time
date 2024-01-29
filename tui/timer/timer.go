package timer

import (
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
}

func New() Model {
	return Model{
		timer: timer.New(),
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
	switch msg.(type) {

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
	return m.timer.View()
}

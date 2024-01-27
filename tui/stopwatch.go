package tui

import (
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// CONSTANTS
// ---------

// The speed at which the clock ticks
const TICK_SPEED = 10 * time.Millisecond

// A tea.Cmd to update the clock every interval
func tick() tea.Cmd {
	return ticker(TICK_SPEED)
}

// MODEL
// -----

// Represents a stopwatch
type stopwatchModel struct {
	running bool          // Whether the stopwatch is running
	start   time.Time     // The time the stopwatch was started
	end     time.Time     // The time the stopwatch was stopped
	elapsed time.Duration // The elapsed time

	keys KeyMap     // Key bindings for the stopwatch view
	help help.Model // Help menu model
}

// Creates a new stopwatch model. Initializes the default key bindings and
// help menu. The stopwatch is "stopped" by default.
func NewStopwatchModel() *stopwatchModel {
	return &stopwatchModel{
		keys: DefaultKeyMap,
		help: help.New(),
	}
}

// INIT, UPDATE, VIEW
// ------------------

// The Init function is called when the program starts and returns a command
// to run immediately. We'll start the stopwatch ticker.
func (m *stopwatchModel) Init() tea.Cmd {
	return tick()
}

// The Update function is called when events happen, like keypresses
func (m *stopwatchModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Window Resize Event
	case tea.WindowSizeMsg:
		// If we set a width on the help menu it can gracefully truncate
		// its view as needed.
		m.help.Width = msg.Width

	// Stopwatch Start
	case msgStart:
		m.start = msg.t
		m.running = true
		m.keys.Space.SetHelp("<spacebar>", "Stop")
		return m, nil

	// Stopwatch Stop
	case msgStop:
		m.end = time.Now()
		m.running = false
		m.keys.Space.SetHelp("<spacebar>", "Start")
		return m, nil

	// Stopwatch Tick
	case msgTick:
		if m.running {
			m.elapsed = time.Since(m.start)
		}
		return m, tick()

	// Key Press
	case tea.KeyMsg:
		switch {

		// Quit
		case key.Matches(msg, DefaultKeyMap.Quit):
			return m, tea.Quit

		// Space
		case key.Matches(msg, DefaultKeyMap.Space):
			if m.running {
				return m, m.Stop
			} else {
				return m, m.Start
			}

		}

	}

	return m, nil
}

// The View function is responsible for rendering the UI.
func (m *stopwatchModel) View() string {
	s := "\n"

	s += m.elapsed.Truncate(TICK_SPEED).String()

	s += "\n\n"

	s += m.help.View(m.keys)

	return s
}

// -------
// COMMAND
// -------

// A command to start the stopwatch
func ShowStopwatch() {
	p := tea.NewProgram(NewStopwatchModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

// MESSAGES & COMMANDS
// -------------------

// A message to start the stopwatch
type msgStart struct{ t time.Time }

// A message to stop the stopwatch
type msgStop struct{}

// A command to start the stopwatch
func (m *stopwatchModel) Start() tea.Msg {
	return msgStart{time.Now()}
}

// A command to stop the stopwatch
func (m *stopwatchModel) Stop() tea.Msg {
	return msgStop{}
}

// KEYBINDINGS
// -----------

// KeyMap is a collection of key bindings
type KeyMap struct {
	Space key.Binding
	Quit  key.Binding
}

var DefaultKeyMap = KeyMap{
	Space: key.NewBinding(key.WithKeys("s", " "), key.WithHelp("<space>", "Start")),
	Quit:  key.NewBinding(key.WithKeys("q", "esc", "ctrl+c"), key.WithHelp("q", "Quit")),
}

// ShortHelp returns a slice of key bindings that are used in the short help
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Space, k.Quit}
}

// FullHelp returns a slice of key bindings that are used in the full help
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Space, k.Quit},
	}
}

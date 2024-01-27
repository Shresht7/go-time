package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// Returns the elapsed time as a string in the format "HH:MM:SS:MS"
func (m *stopwatchModel) Elapsed() string {
	hours := int(m.elapsed.Hours()) % 24
	minutes := int(m.elapsed.Minutes()) % 60
	seconds := int(m.elapsed.Seconds()) % 60
	millisecond := int(m.elapsed.Milliseconds() % 1000)
	return fmt.Sprintf("%02dh : %02dm : %02ds : %03dms", hours, minutes, seconds, millisecond)
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
		m.keys.Space.SetHelp("<spacebar>", "stop")
		return m, nil

	// Stopwatch Stop
	case msgStop:
		m.end = time.Now()
		m.running = false
		m.keys.Space.SetHelp("<spacebar>", "start")
		return m, nil

	// Stopwatch Resume
	case msgResume:
		m.start = time.Now().Add(-m.elapsed)
		m.running = true
		m.keys.Space.SetHelp("<spacebar>", "stop")
		return m, nil

	// Stopwatch Reset
	case msgReset:
		m.elapsed = 0
		m.running = false
		m.keys.Space.SetHelp("<spacebar>", "start")
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
				return m, m.Resume
			}

		// R
		case key.Matches(msg, DefaultKeyMap.R):
			return m, m.Reset

		}

	}

	return m, nil
}

var styles = lipgloss.NewStyle()

// The View function is responsible for rendering the UI.
func (m *stopwatchModel) View() string {
	s := "\n"

	time := m.Elapsed()
	s += styles.Render(time) + "\n\n"

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

// A message to resume the stopwatch
type msgResume struct{}

// A message to reset the stopwatch
type msgReset struct{}

// A command to start the stopwatch
func (m *stopwatchModel) Start() tea.Msg {
	return msgStart{time.Now()}
}

// A command to stop the stopwatch
func (m *stopwatchModel) Stop() tea.Msg {
	return msgStop{}
}

// A command to resume the stopwatch
func (m *stopwatchModel) Resume() tea.Msg {
	return msgResume{}
}

// A command to reset the stopwatch
func (m *stopwatchModel) Reset() tea.Msg {
	return msgReset{}
}

// KEYBINDINGS
// -----------

// KeyMap is a collection of key bindings
type KeyMap struct {
	R     key.Binding
	Space key.Binding
	Quit  key.Binding
}

var DefaultKeyMap = KeyMap{
	Space: key.NewBinding(key.WithKeys("s", " "), key.WithHelp("<spacebar>", "start")),
	R:     key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "reset")),
	Quit:  key.NewBinding(key.WithKeys("q", "esc", "ctrl+c"), key.WithHelp("q", "quit")),
}

// ShortHelp returns a slice of key bindings that are used in the short help
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.R, k.Space}
}

// FullHelp returns a slice of key bindings that are used in the full help
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.R, k.Space},
	}
}

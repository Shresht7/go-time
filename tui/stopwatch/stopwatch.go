package stopwatch

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// MODEL
// -----

// Returns the elapsed time as a string in the format "HH:MM:SS:MS"
func formatElapsed(elapsed time.Duration) string {
	hours := int(elapsed.Hours()) % 24
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60
	millisecond := int(elapsed.Milliseconds() % 1000)
	return fmt.Sprintf("%02dh : %02dm : %02ds : %03dms", hours, minutes, seconds, millisecond)
}

// INIT, UPDATE, VIEW
// ------------------

func Filter(s ...string) []string {
	var filtered []string
	for _, v := range s {
		if v != "" {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

// -------
// COMMAND
// -------

// A command to start the stopwatch
func Run() {
	p := tea.NewProgram(newStopwatchModel())
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

type msgLap struct{}

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

func (m *stopwatchModel) Lap() tea.Msg {
	return msgLap{}
}

// KEYBINDINGS
// -----------

// KeyMap is a collection of key bindings
type KeyMap struct {
	R     key.Binding
	Space key.Binding
	Enter key.Binding
	Quit  key.Binding
}

var DefaultKeyMap = KeyMap{
	Space: key.NewBinding(key.WithKeys("s", " "), key.WithHelp("<spacebar>", "start")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("<enter>", "lap")),
	R:     key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "reset")),
	Quit:  key.NewBinding(key.WithKeys("q", "esc", "ctrl+c"), key.WithHelp("q", "quit")),
}

// ShortHelp returns a slice of key bindings that are used in the short help
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.R, k.Enter, k.Space}
}

// FullHelp returns a slice of key bindings that are used in the full help
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Quit, k.R, k.Enter, k.Space},
	}
}

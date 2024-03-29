package stopwatch

import "github.com/charmbracelet/bubbles/key"

// KEYBINDINGS
// -----------

// KeyMap is a collection of key bindings
type KeyMap struct {
	R     key.Binding
	Space key.Binding
	Enter key.Binding
	Quit  key.Binding
}

// The default key bindings for the stopwatch
var DefaultKeyMap = KeyMap{
	Space: key.NewBinding(key.WithKeys("s", " "), key.WithHelp("spacebar", "start")),
	Enter: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "lap")),
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

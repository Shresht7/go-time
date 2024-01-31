package timer

import "github.com/charmbracelet/bubbles/key"

// KEYBINDINGS
// -----------

// KeyMap is a collection of keybindings
type KeyMap struct {
	Tab  key.Binding
	Quit key.Binding
}

// The default keybindings for the timer program
var DefaultKeyMap = KeyMap{
	Tab:  key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "switch")),
	Quit: key.NewBinding(key.WithKeys("q", "esc", "ctrl+c"), key.WithHelp("q", "quit")),
}

// ShortHelp returns a slice of keybindings that are used in the short help
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Tab, k.Quit}
}

// FullHelp returns a slice of keybindings that are used in the full help
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Tab, k.Quit},
	}
}

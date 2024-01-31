package timer

import "github.com/charmbracelet/bubbles/key"

// KEYBINDINGS
// -----------

// KeyMap is a collection of keybindings
type KeyMap struct {
	Space key.Binding
	Reset key.Binding
}

// The default keybindings for the timer subcomponent
var DefaultKeyMap = KeyMap{
	Space: key.NewBinding(key.WithKeys("s", " "), key.WithHelp("spacebar", "start")),
	Reset: key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "reset")),
}

// ShortHelp returns a slice of keybindings that are used in the short help
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Space, k.Reset}
}

// FullHelp returns a slice of keybindings that are used in the full help
func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Space, k.Reset},
	}
}

package list

import "github.com/charmbracelet/bubbles/key"

// KEYBINDINGS
// -----------

// KeyModel holds the keybindings for the list component
type KeyModel struct {
	Select   key.Binding
	Previous key.Binding
	Next     key.Binding
}

// DefaultKeyModel is the default keybindings for the list component
var DefaultKeyModel = KeyModel{
	Select:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	Previous: key.NewBinding(key.WithKeys("up", "k"), key.WithHelp("up/k", "previous")),
	Next:     key.NewBinding(key.WithKeys("down", "j"), key.WithHelp("down/j", "next")),
}

// AltKeyModel is an alternate keybindings for the list component
var AltKeyModel = KeyModel{
	Select:   key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "select")),
	Previous: key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("left/h", "previous")),
	Next:     key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("right/l", "next")),
}

// ShortHelp returns a slice of keybindings that are used in the short help
func (m KeyModel) ShortHelp() []key.Binding {
	return []key.Binding{m.Previous, m.Next, m.Select}
}

// FullHelp returns a slice of keybindings that are used in the full help
func (m KeyModel) FullHelp() [][]key.Binding {
	return [][]key.Binding{{m.Previous, m.Next, m.Select}}
}

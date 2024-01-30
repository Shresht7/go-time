package calendar

import "github.com/charmbracelet/bubbles/key"

type KeyMaps struct {
	Left  key.Binding
	Right key.Binding
	Quit  key.Binding
}

var DefaultKeyMaps = KeyMaps{
	Left:  key.NewBinding(key.WithKeys("left", "h"), key.WithHelp("←/h", "previous")),
	Right: key.NewBinding(key.WithKeys("right", "l"), key.WithHelp("→/l", "next")),
	Quit:  key.NewBinding(key.WithKeys("q", "esc", "ctrl+c"), key.WithHelp("q/esc/ctrl+c", "quit")),
}

func (k KeyMaps) ShortHelp() []key.Binding {
	return []key.Binding{k.Left, k.Right, k.Quit}
}

func (k KeyMaps) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Left, k.Right}, {k.Quit}}
}

package timer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"
)

// MODEL
// -----

type Model struct {
	Running   bool // Whether the timer is running
	remaining int  // Remaining time in seconds
	preset    int  // The preset time in seconds

	isFocused bool // Whether the timer is focused

	spinner spinner.Model // Spinner model
	keys    KeyMap        // Keybindings model
	help    help.Model    // Help model
}

func New() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	return Model{
		remaining: 300,
		preset:    300,
		spinner:   s,
		keys:      DefaultKeyMap,
		help:      help.New(),
	}
}

func (m *Model) Start() {
	m.Running = true
	m.keys.Space.SetHelp("spacebar", "pause")
}

func (m *Model) Stop() {
	m.Running = false
	m.keys.Space.SetHelp("spacebar", "resume")
}

func (m *Model) Set(t string) {
	m.Stop()
	switch t {
	case "1 Minute":
		m.preset = 60
	case "5 Minutes":
		m.preset = 300
	case "10 Minutes":
		m.preset = 600
	case "15 Minutes":
		m.preset = 900
	case "30 Minutes":
		m.preset = 1800
	case "60 Minutes":
		m.preset = 3600
	default:
		m.preset = 60
	}
	m.remaining = m.preset
}

func (m *Model) Reset() {
	m.remaining = m.preset
}

// Set the focus state of the timer
func (m *Model) SetFocused(f bool) {
	m.isFocused = f
}

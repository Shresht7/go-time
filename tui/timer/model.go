package timer

import (
	"github.com/charmbracelet/bubbles/help"

	"github.com/Shresht7/go-time/tui/timer/components/list"
	"github.com/Shresht7/go-time/tui/timer/components/timer"
)

// FOCUS
// -----

// Indicates which model is currently focused
type focused int

// An enumeration of the models that can be focused
const (
	focusTimer focused = iota
	focusList
)

// MODEL
// -----

// A model representing the state of the timer program
type Model struct {
	timer timer.Model // The timer model
	list  list.Model  // The list model

	focused focused // Indicates which model is currently focused

	keys KeyMap     // Key bindings model
	help help.Model // Help menu model
}

// Creates a new timer program model
func New() Model {
	m := Model{
		timer: timer.New(),
		// TODO: make this configurable
		list: list.New(
			"00:01:00",
			"00:05:00",
			"00:10:00",
			"00:15:00",
			"00:30:00",
			"00:45:00",
			"00:60:00",
		),
		keys: DefaultKeyMap,
		help: help.New(),
	}
	m.list.SetPrompt("Select a preset:\n") // Set the prompt message for the list
	m.SetFocus(focusTimer)                 // focus the timer by default
	return m
}

// Set the focus to the given model enumeration
func (m *Model) SetFocus(f focused) {
	m.focused = f

	switch f {

	case focusTimer:
		m.timer.SetFocused(true)
		m.list.SetFocused(false)
		m.keys.Tab.SetHelp("tab", "presets")

	case focusList:
		m.timer.SetFocused(false)
		m.list.SetFocused(true)
		m.keys.Tab.SetHelp("tab", "timer")

	}
}

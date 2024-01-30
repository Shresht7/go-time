package list

import "github.com/charmbracelet/bubbles/help"

// MODEL
// -----

// Holds the state of the list component
type Model struct {
	items    []string // items in the list
	selected int      // index of the selected item
	prompt   string   // prompt to display above the list

	isFocused bool // whether the list is focused

	keys KeyModel   // keybindings model
	help help.Model // help model
}

// Instantiates a new list model
func New(items ...string) Model {
	return Model{
		items:     items,
		selected:  0,
		isFocused: false,
		keys:      DefaultKeyModel,
		help:      help.New(),
	}
}

// Get the item at the given index
func (m Model) At(index int) string {
	return m.items[index]
}

// Set the prompt
func (m *Model) SetPrompt(p string) {
	m.prompt = p
}

// Move the selection to the previous item
func (m *Model) Previous() {
	if m.selected > 0 {
		m.selected--
	}
}

// Move the selection to the next item
func (m *Model) Next() {
	if m.selected < len(m.items)-1 {
		m.selected++
	}
}

// Set the focused state of the list
func (m *Model) SetFocused(f bool) {
	m.isFocused = f
}

// Returns the help menu for the list component
func (m *Model) ViewHelp() string {
	return m.help.View(m.keys)
}

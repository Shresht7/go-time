package list

import "github.com/charmbracelet/bubbles/help"

// MODEL
// -----

// Holds the state of the list component
type Model struct {
	items    []string // items in the list
	selected int      // index of the selected item

	keys KeyModel   // keybindings model
	help help.Model // help model
}

// Instantiates a new list model
func New(items ...string) Model {
	return Model{
		items:    items,
		selected: 0,
		keys:     DefaultKeyModel,
		help:     help.New(),
	}
}

// Returns the list of items
func (m Model) Items() []string {
	return m.items
}

// Returns the index of the selected item
func (m Model) Selected() int {
	return m.selected
}

// Returns the selected item
func (m Model) SelectedItem() string {
	return m.items[m.selected]
}

// Returns the length of the list
func (m Model) Len() int {
	return len(m.items)
}

// Move the selection to the previous item
func (m Model) Previous() Model {
	if m.selected > 0 {
		m.selected--
	}
	return m
}

// Move the selection to the next item
func (m Model) Next() Model {
	if m.selected < len(m.items)-1 {
		m.selected++
	}
	return m
}

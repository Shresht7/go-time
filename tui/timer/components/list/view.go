package list

// VIEW
// ----

// Returns the view for the list component
func (m Model) View() string {
	var s string
	for i, item := range m.items {
		if i == m.selected {
			s += " > " + item + "\n"
		} else {
			s += "   " + item + "\n"
		}
	}
	return s
}

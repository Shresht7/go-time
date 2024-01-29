package clock

// VIEW
// ----

// The View function is called every time the model is updated
func (c *model) View() string {
	s := "\n" // start with a newline for padding

	s += c.Icon() // add the icon for the current time of day

	// add the time in the format "HH:MM:SS"
	hours, minutes, seconds := c.formatTime()
	s += "  " + hours + ":" + minutes + ":" + seconds + "\t"

	// add the date in the format "Saturday, 27 January 2024"
	s += c.t.Format("Monday, 2 January 2006")

	// add a newline and some padding at the bottom
	s += "\n\n"

	return s
}

package laps

import "time"

// MODEL
// -----

// Represents a list of lap times
type Model struct {
	laps []lap // The collection of lap times
}

// Creates a new lap times model
func New() Model {
	return Model{}
}

// Adds a lap time to the list
func (m *Model) Add(start, end time.Time) {
	m.laps = append(m.laps, lap{start, end})
}

// Clears the list of lap times
func (m *Model) Reset() {
	m.laps = []lap{}
}

// Returns the number of lap times in the list
func (m Model) Len() int {
	return len(m.laps)
}

// Returns the last lap time in the list or a fallback time if the list is empty
func (m Model) LastOr(fallback time.Time) time.Time {
	if m.Len() == 0 {
		return fallback
	}
	return m.laps[m.Len()-1].end
}

// Returns the shortest lap time in the list
func (m Model) Shortest() int {
	shortest := 0
	for i, lap := range m.laps {
		if lap.duration() < m.laps[shortest].duration() {
			shortest = i
		}
	}
	return shortest
}

// Returns the longest lap time in the list
func (m Model) Longest() int {
	longest := 0
	for i, lap := range m.laps {
		if lap.duration() > m.laps[longest].duration() {
			longest = i
		}
	}
	return longest
}

// Returns the total time of all lap times in the list
func (m Model) Sum(i int) time.Duration {
	var total time.Duration
	for _, lap := range m.laps[:i+1] {
		total += lap.duration()
	}
	return total
}

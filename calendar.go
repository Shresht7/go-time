package main

import (
	"fmt"
	"time"
)

var WEEKDAYS = []string{"M", "T", "W", "T", "F", "S", "S"}

type Calendar struct {
	date  string
	month time.Month
	year  int
	grid  [6][]string
}

//	Creates a new Calendar
func NewCalendar(now time.Time) Calendar {
	date := fmt.Sprintf("%d", now.Day())
	month := now.Month()
	year := now.Year()

	grid := [6][]string{}
	grid[0] = WEEKDAYS
	for i := 1; i < len(grid); i++ {
		grid[i] = []string{}
		for j := 1; j < 8; j++ {
			grid[i] = append(grid[i], fmt.Sprintf("%d%d", i, j))
		}
	}

	return Calendar{date: date, month: month, year: year, grid: grid}
}

//	Print the calendar to the screen
func (cal *Calendar) show() {
	fmt.Printf("\n%11s %d\n", cal.month, cal.year)
	for _, row := range cal.grid {
		for r, date := range row {
			str := ""
			if r == 6 {
				str = fmt.Sprintf("\u001b[31m%2s \u001b[39m", date)
			} else {
				str = fmt.Sprintf("%2s ", date)
			}
			if date == cal.date {
				str = fmt.Sprintf("\u001b[7m%2s\u001b[27m ", date)
			}
			fmt.Print(str)
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

//	Show calendar command
func showCalendar(now time.Time) {
	calendar := NewCalendar(now)
	calendar.show()
}

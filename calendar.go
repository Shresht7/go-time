package main

import (
	"fmt"
	"time"
)

type Calendar struct {
	month time.Month
	year  int
	grid  [6][7]string
}

func NewCalendar(now time.Time) Calendar {
	month := now.Month()
	year := now.Year()

	grid := [6][7]string{}
	grid[0] = [7]string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	for i := 1; i < len(grid); i++ {
		grid[i] = [7]string{"1", "2", "3", "4", "5", "6", "7"}
	}
	return Calendar{month: month, year: year, grid: grid}
}

func (cal *Calendar) show() {
	fmt.Printf("\n%11s %d\n", cal.month, cal.year)
	for _, row := range cal.grid {
		for r, el := range row {
			if r == 6 {
				fmt.Printf("\u001b[31m%2s\u001b[39m", el)
			} else {
				fmt.Printf("%2s ", el)
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func showCalendar(now time.Time) {
	calendar := NewCalendar(now)
	calendar.show()
}

package calendar

import (
	"fmt"
	"strings"
	"time"

	"github.com/Shresht7/go-cli-tools/ansi/colors"
	"github.com/Shresht7/go-cli-tools/ansi/styles"
)

// MONTH
// -----

// TODO: This is utter chaos and needs a complete rewrite!

var WEEKDAYS = []string{"M", "T", "W", "T", "F", "S", "S"}

// Renders a calender string
func Render(t time.Time) string {
	var sb strings.Builder

	// Create calendar grid
	grid := CreateCalendarGrid(t)

	// Add month and year
	sb.WriteString(MonthAndYear(t))

	// Add weekday headers
	sb.WriteString(WeekdayHeaders())

	// Add calendar grid
	for _, row := range grid {
		sb.WriteString(strings.Join(row, " "))
		sb.WriteString("\n")
	}

	// Add empty line
	sb.WriteString("\n")

	// Build and return string
	return sb.String()
}

func MonthAndYear(t time.Time) string {
	return fmt.Sprintf(styles.Bold("\n%11s %d\n\n"), t.Month(), t.Year())
}

func WeekdayHeaders() string {
	s := ""
	for c := 0; c < 7; c++ {
		s += fmt.Sprintf(styles.Bold("%2s"), WEEKDAYS[c]) + " "
	}
	return s
}

func CreateCalendarGrid(t time.Time) [5][]string {
	//	Grid to hold dates
	grid := [5][]string{}

	// Get the first day of the month
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).Weekday()

	//	Add dates to calendar grid
	for r := 1; r < len(grid); r++ { // We start from 1 because the first row is for weekday headers
		grid[r] = []string{}
		for c := 1; c <= 7; c++ {

			//	Get date of month
			d := time.Date(t.Year(), t.Month(), ((r-1)*7)+c+1-int(firstDay), 0, 0, 0, 0, t.Location())
			str := fmt.Sprintf("%2d", d.Day())

			//	If date is of previous or next month then make it faint
			if d.Month() != t.Month() {
				str = styles.Faint(str)
			}

			//	If date is a sunday, color it red
			if c == 7 {
				str = colors.Red(str)
			}

			//	If today, invert color
			if d.Day() == time.Now().Day() && d.Month() == time.Now().Month() && d.Year() == time.Now().Year() {
				str = styles.Inverse(str)
			}

			grid[r] = append(grid[r], str)
		}
	}

	return grid
}

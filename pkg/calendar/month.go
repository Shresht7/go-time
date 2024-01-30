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

// Creates a new Calendar
func NewMonth(t time.Time) string {
	// Create calendar grid
	grid := CreateCalendarGrid(t)

	//	Convert grid to string
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(styles.Bold("\n%11s %d\n\n"), t.Month(), t.Year()))

	//	Add weekday headers to the first row
	for c := 0; c < 7; c++ {
		sb.WriteString(fmt.Sprintf(styles.Bold("%2s"), WEEKDAYS[c]))
		sb.WriteString(" ")
	}

	for _, row := range grid {
		sb.WriteString(strings.Join(row, " "))
		sb.WriteString("\n")
	}

	sb.WriteString("\n")

	return sb.String()
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
			if d.Day() == t.Day() {
				str = styles.Inverse(str)
			}

			grid[r] = append(grid[r], str)
		}
	}

	return grid
}

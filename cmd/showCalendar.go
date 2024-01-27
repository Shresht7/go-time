package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Shresht7/go-cli-tools/ansi/colors"
	"github.com/Shresht7/go-cli-tools/ansi/styles"
)

var WEEKDAYS = []string{"M", "T", "W", "T", "F", "S", "S"}

type Calendar struct {
	date  string
	month time.Month
	year  int
	grid  [6][]string
}

// Creates a new Calendar
func NewCalendar(now time.Time) Calendar {
	date := strconv.Itoa(now.Day())
	month := now.Month()
	year := now.Year()

	//	Grid to hold dates
	grid := [6][]string{}

	//	Add weekday headers
	for c := 0; c < 7; c++ {
		grid[0] = append(grid[0], fmt.Sprintf("%2s", WEEKDAYS[c]))
	}

	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, now.Location()).Day() //	First day of the month

	//	Add dates to calendar grid
	for r := 1; r < len(grid); r++ {
		grid[r] = []string{}
		for c := 1; c < 8; c++ {

			//	Get date of month
			d := time.Date(year, month, c+((r-1)*7)-firstDay, 0, 0, 0, 0, now.Location())
			str := fmt.Sprintf("%2d", d.Day())

			//	If date is of previous or next month then prefix with ~. Used later for special formatting
			if d.Month() != month {
				str = fmt.Sprintf("~%s", str)
			}

			grid[r] = append(grid[r], str)
		}
	}

	return Calendar{date: date, month: month, year: year, grid: grid}
}

// Print the calendar to the screen
func (cal *Calendar) show() {

	//	Print month and calendar
	fmt.Printf("\n%11s %d\n", cal.month, cal.year)

	for _, row := range cal.grid {
		for c, date := range row {

			str := date

			//	If prefixed with ~ (date not of this month), reduce brightness
			if strings.HasPrefix(str, "~") {
				str = str[1:]
				str = styles.Faint(str)
			}

			//	If sunday column, color it red
			if c == 6 {
				str = colors.Red(fmt.Sprintf("%2s", str))
			}

			//	If today, invert color
			if strings.TrimSpace(date) == strings.TrimSpace(cal.date) {
				str = styles.Inverse(fmt.Sprintf("%2s", str))
			}

			//	Spacing
			fmt.Print(str + " ")
		}

		fmt.Print("\n")

	}

	fmt.Println() //	Empty line

}

// Show calendar command
func ShowCalendar(now time.Time) {
	calendar := NewCalendar(now)
	calendar.show()
}

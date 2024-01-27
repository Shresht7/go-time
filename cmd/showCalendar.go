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

type calendarMonth struct {
	t    time.Time
	grid [6][]string
}

// Creates a new Calendar
func NewCalendarMonth(t time.Time) *calendarMonth {

	//	Grid to hold dates
	grid := [6][]string{}

	//	Add weekday headers
	for c := 0; c < 7; c++ {
		grid[0] = append(grid[0], fmt.Sprintf("%2s", WEEKDAYS[c]))
	}

	// Get the first day of the month
	firstDay := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location()).Day() //	First day of the month

	//	Add dates to calendar grid
	for r := 1; r < len(grid); r++ {
		grid[r] = []string{}
		for c := 1; c < 8; c++ {

			//	Get date of month
			d := time.Date(t.Year(), t.Month(), c+((r-1)*7)-firstDay, 0, 0, 0, 0, t.Location())
			str := fmt.Sprintf("%2d", d.Day())

			//	If date is of previous or next month then prefix with ~. Used later for special formatting
			if d.Month() != t.Month() {
				str = fmt.Sprintf("~%s", str)
			}

			grid[r] = append(grid[r], str)
		}
	}

	return &calendarMonth{t, grid}
}

// Print the calendar to the screen
func (cal *calendarMonth) show() {

	//	Print month and calendar
	fmt.Printf(styles.Bold("\n%11s %d\n\n"), cal.t.Month(), cal.t.Year())

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
			if strings.TrimSpace(date) == strings.TrimSpace(strconv.Itoa(cal.t.Day())) {
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
	calendar := NewCalendarMonth(now)
	calendar.show()
}

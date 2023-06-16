package calendar

import (
	"flag"
	"fmt"
	"nepdate/internal"
	"nepdate/internal/utils"
	"os"
	"time"
)

var calendarFlagSet = flag.NewFlagSet("cal", flag.ExitOnError)
var yearFlag = calendarFlagSet.Bool("y", false, "nepdate cal -y")

const ERROR_PARSING_FLAG = "Calendar: Error while parsing flags %s"

func Calendar() error {
	if err := calendarFlagSet.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, ERROR_PARSING_FLAG, err.Error())
		os.Exit(1)
	}
	currentDate, err := internal.GetCurrentDate()
	if err != nil {
		return fmt.Errorf("Calendar: %v", err)
	}

	monthsToPrint := []int{currentDate.Month}
	if *yearFlag {
		monthsToPrint = utils.GetMonths()
	}
	err = PrintMonthsCalendar(currentDate.Year, monthsToPrint, currentDate)
	if err != nil {
		return fmt.Errorf("Calendar: %v", err)
	}
	return nil
}

// currentDate is for highlighting
func PrintMonthsCalendar(year int, months []int, currentDate internal.NepaliDate) error {
	for _, m := range months {
		// Print title
		fmt.Printf("%s%s %d\n", "     ", internal.GetMonthName(m), currentDate.Year)

		// Print weekdays
		weekDays := utils.GetWeekDays()
		for _, weekDay := range weekDays {
			fmt.Printf("%s ", utils.WeekdayShortname(weekDay))
		}
		fmt.Println()

		// Start printing from 1st of month to last
		monthStart, err := internal.GetNepaliDate(year, m, 1)
		if err != nil {
			return fmt.Errorf("PrintMonthCalendar: %v", err)
		}

		daysInMonth := internal.GetCurrentYearInfo().DaysInMonths
		startWeekDay := time.Unix(monthStart.Timestamp, 0).Weekday()
		// Print spaces
		for i := 0; i < int(startWeekDay); i++ {
			fmt.Printf("   ")
		}
		// Print months
		for i := 0; i < daysInMonth[m]; i++ {
			if i+1 == currentDate.Date && m == currentDate.Month && currentDate.Year == year {
				fmt.Printf("\x1b[30;47m%2d\x1b[0m ", i+1)
			} else {
				fmt.Printf("%2d ", i+1)
			}
			startWeekDay += 1
			if startWeekDay == 7 {
				fmt.Printf("\n")
				startWeekDay = 0
			}
		}
		fmt.Printf("\n")
	}
	return nil
}

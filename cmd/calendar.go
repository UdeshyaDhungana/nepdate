package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/UdeshyaDhungana/nepdate/internal"
	"github.com/spf13/cobra"
)

var calYear bool

var calCmd = &cobra.Command{
	Use:   "cal",
	Short: "Print calendar info",
	Run: func(cmd *cobra.Command, args []string) {
		today := internal.MustConvertAdTimeToBsDate(time.Now())
		sakrati := today.GetSakrati()
		calendarWidth := 20 // 2 spaces for each day of the week + 6 gaps

		// Print the title
		calendarTitle := fmt.Sprintf("%v %d", today.Month, today.Year)
		calendarTitleOffset := (calendarWidth - len(calendarTitle)) / 2
		fmt.Println(strings.Repeat(" ", calendarTitleOffset), calendarTitle)

		// Print the weekday names (Su - Sa)
		for weekday := time.Sunday; weekday <= time.Saturday; weekday++ {
			fmt.Printf("%s", weekday.String()[:2])
			if weekday == time.Saturday {
				fmt.Printf("\n")
			} else {
				fmt.Printf(" ")
			}
		}

		// Print the days of the month
		daysCountInThisMonth := internal.GetConfig().Years[today.Year].DaysInMonths[today.Month]
		sakratiWeekday := int(sakrati.Weekday())
		weeksCountInThisMonth := 1 + (sakratiWeekday+daysCountInThisMonth)/7

		for weekIdx := range weeksCountInThisMonth {
			sundayMiti := 1 + 7*weekIdx - sakratiWeekday
			saturdayMiti := sundayMiti + 6

			var thisWeekDays []string
			for miti := sundayMiti; miti <= saturdayMiti; miti++ {
				if miti <= 0 || miti > daysCountInThisMonth {
					thisWeekDays = append(thisWeekDays, "  ")
					continue
				}

				if miti == today.Day {
					thisWeekDays = append(thisWeekDays, fmt.Sprintf("\x1b[30;47m%02d\x1b[0m", miti))
					continue
				}

				thisWeekDays = append(thisWeekDays, fmt.Sprintf("%02d", miti))
			}

			fmt.Printf("%s\n", strings.Join(thisWeekDays, " "))
		}
	},
}

func init() {
	// calCmd.Flags().BoolVarP(&calYear, "year", "y", false, "show the whole year instead of just this month")
	rootCmd.AddCommand(calCmd)
}

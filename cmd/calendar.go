package cmd

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/UdeshyaDhungana/nepdate/internal"
	"github.com/spf13/cobra"
)

var calYear, calExtended bool

const calendarWidth = 20 // 7 days × 2 chars + 6 gaps

// ansiEscape matches ANSI color escape sequences so we can measure visible length.
var ansiEscape = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// padToWidth pads s with trailing spaces until its visible width equals width.
func padToWidth(s string, width int) string {
	visibleLen := len(ansiEscape.ReplaceAllString(s, "")) // strip ANSI codes to get printable length
	if pad := width - visibleLen; pad > 0 {
		return s + strings.Repeat(" ", pad)
	}
	return s
}

// renderMonth builds a single month's calendar as a slice of lines.
// Pass today non-nil to highlight that day; pass nil for no highlight.
// Pass showYear false to omit the year from the title (e.g. when the year is already shown above).
func renderMonth(year int, month internal.BSMonth, today *internal.BSDate, showYear bool) []string {
	var lines []string

	// Title: centered month name, with year only when not already shown above
	title := fmt.Sprintf("%v", month)
	if showYear {
		title = fmt.Sprintf("%v %d", month, year)
	}
	titleOffset := (calendarWidth - len(title)) / 2
	if titleOffset < 0 {
		titleOffset = 0
	}
	lines = append(lines, strings.Repeat(" ", titleOffset)+title)

	// Weekday header
	var weekdays []string
	for wd := time.Sunday; wd <= time.Saturday; wd++ {
		weekdays = append(weekdays, wd.String()[:2])
	}
	lines = append(lines, strings.Join(weekdays, " "))

	// Days
	sakrati := internal.BSDate{Year: year, Month: month, Day: 1}
	sakratiWeekday := int(sakrati.Weekday())
	daysCount := internal.GetConfig().Years[year].DaysInMonths[month]
	weeksCount := 1 + (sakratiWeekday+daysCount)/7

	for weekIdx := range weeksCount {
		sundayMiti := 1 + 7*weekIdx - sakratiWeekday
		saturdayMiti := sundayMiti + 6

		var dayStrs []string
		for miti := sundayMiti; miti <= saturdayMiti; miti++ {
			if miti <= 0 || miti > daysCount {
				dayStrs = append(dayStrs, "  ")
				continue
			}
			if today != nil && today.Year == year && today.Month == month && today.Day == miti {
				dayStrs = append(dayStrs, fmt.Sprintf("\x1b[30;47m%02d\x1b[0m", miti))
				continue
			}
			dayStrs = append(dayStrs, fmt.Sprintf("%02d", miti))
		}
		lines = append(lines, strings.Join(dayStrs, " "))
	}

	return lines
}

// printYearCalendar prints all 12 months laid out in a 3-column × 4-row grid.
func printYearCalendar(today internal.BSDate) {
	const cols = 3
	const colGap = "  "

	// Centered year header
	yearStr := fmt.Sprintf("%d", today.Year)
	totalWidth := cols*calendarWidth + (cols-1)*len(colGap)
	yearOffset := (totalWidth - len(yearStr)) / 2
	fmt.Println(strings.Repeat(" ", yearOffset) + yearStr)
	fmt.Println()

	for row := range 4 {
		// Render the 3 months in this row
		var blocks [][]string
		maxLines := 0
		for col := range cols {
			month := internal.BSMonth(row*cols + col)
			block := renderMonth(today.Year, month, &today, false)
			blocks = append(blocks, block)
			if len(block) > maxLines {
				maxLines = len(block)
			}
		}

		// Pad all blocks to the same height
		for i := range blocks {
			for len(blocks[i]) < maxLines {
				blocks[i] = append(blocks[i], "")
			}
		}

		// Print each line across the 3 columns
		for lineIdx := range maxLines {
			var parts []string
			for col := range cols {
				parts = append(parts, padToWidth(blocks[col][lineIdx], calendarWidth))
			}
			fmt.Println(strings.Join(parts, colGap))
		}
		fmt.Println()
	}
}

var calCmd = &cobra.Command{
	Use:   "cal",
	Short: "Print calendar info",
	Run: func(cmd *cobra.Command, args []string) {
		today := internal.MustConvertAdTimeToBsDate(time.Now())
		if calYear {
			printYearCalendar(today)
		} else if calExtended {
			printExtendedMonth(today.Year, today.Month, &today)
		} else {
			for _, line := range renderMonth(today.Year, today.Month, &today, true) {
				fmt.Println(line)
			}
		}
	},
}

func init() {
	calCmd.Flags().BoolVarP(&calYear, "year", "y", false, "show the whole year's calendar")
	calCmd.Flags().BoolVarP(&calExtended, "extended", "e", false, "show a large grid calendar for the month, with A.D. dates")
	calCmd.MarkFlagsMutuallyExclusive("year", "extended")
	rootCmd.AddCommand(calCmd)
}

package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/UdeshyaDhungana/nepdate/internal"
)

const (
	extCellWidth = 9 // inner width of a cell
	extCellRows  = 1 // inner height of a cell
)

var subscriptDigits = strings.NewReplacer(
	"0", "₀", "1", "₁", "2", "₂", "3", "₃", "4", "₄",
	"5", "₅", "6", "₆", "7", "₇", "8", "₈", "9", "₉",
)

// center pads s with spaces on both sides to reach width (s is plain text).
func center(s string, width int) string {
	pad := width - len([]rune(s))
	if pad <= 0 {
		return s
	}
	left := pad / 2
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", pad-left)
}

// renderExtCell returns the inner lines of one day cell: the B.S. day with the A.D. date in subscript.
func renderExtCell(day int, ad time.Time, highlight bool) []string {
	english := subscriptDigits.Replace(fmt.Sprintf("%d", ad.Day()))
	if ad.Day() == 1 {
		english = ad.Format("Jan") + english
	}

	lines := make([]string, extCellRows)
	for i := range lines {
		lines[i] = strings.Repeat(" ", extCellWidth)
	}
	lines[extCellRows/2] = center(fmt.Sprintf("%d %s", day, english), extCellWidth)

	if highlight {
		for i, l := range lines {
			lines[i] = "\x1b[30;47m" + l + "\x1b[0m"
		}
	}
	return lines
}

// printExtendedMonth prints a large grid calendar for the given month.
func printExtendedMonth(year int, month internal.BSMonth, today *internal.BSDate) {
	gridWidth := 7*extCellWidth + 8
	title := fmt.Sprintf("%v %d", month, year)
	fmt.Println(center(title, gridWidth))

	border := func(l, m, r string) string {
		seg := strings.Repeat("─", extCellWidth)
		return l + strings.Repeat(seg+m, 6) + seg + r
	}

	fmt.Println(border("┌", "┬", "┐"))
	var heads []string
	for wd := time.Sunday; wd <= time.Saturday; wd++ {
		heads = append(heads, center(wd.String()[:3], extCellWidth))
	}
	fmt.Println("│" + strings.Join(heads, "│") + "│")
	fmt.Println(border("├", "┼", "┤"))

	sakratiWeekday := int(internal.BSDate{Year: year, Month: month, Day: 1}.Weekday())
	daysCount := internal.GetConfig().Years[year].DaysInMonths[month]
	weeksCount := 1 + (sakratiWeekday+daysCount-1)/7
	blank := strings.Repeat(" ", extCellWidth)

	for week := range weeksCount {
		cells := make([][]string, 7)
		for col := range 7 {
			miti := 1 + 7*week + col - sakratiWeekday
			if miti <= 0 || miti > daysCount {
				cells[col] = make([]string, extCellRows)
				for i := range cells[col] {
					cells[col][i] = blank
				}
				continue
			}
			ad := internal.MustConvertBsDateToAdTime(internal.BSDate{Year: year, Month: month, Day: miti})
			isToday := today != nil && today.Year == year && today.Month == month && today.Day == miti
			cells[col] = renderExtCell(miti, ad, isToday)
		}
		for i := range extCellRows {
			var row []string
			for col := range 7 {
				row = append(row, cells[col][i])
			}
			fmt.Println("│" + strings.Join(row, "│") + "│")
		}
		if week < weeksCount-1 {
			fmt.Println(border("├", "┼", "┤"))
		}
	}
	fmt.Println(border("└", "┴", "┘"))
}

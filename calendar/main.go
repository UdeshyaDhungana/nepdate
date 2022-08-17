package calendar

import (
	"fmt"
	"nepdate/utils"
	"strings"

	"github.com/gocolly/colly"
)

const URL = "https://nepalicalendar.rat32.com"
const COMMON_SELECTOR = "html>body>div#page>div#content>div#leftMiddle>div#father>div#cover"
const MONTH_SELECTOR = "div#monthtitle>h1#yren"
const DATE_SELECTOR = "div#main"

func Calendar() {
	c := colly.NewCollector()
	var monthTitle string
	var gaps int8
	var dontSkipgaps bool
	var endDate int8
	var currentDate int8

	// On Error
	c.OnError(func(r *colly.Response, err error) {
		panic("Network connection error.")
	})

	// On HTML Response
	c.OnHTML(COMMON_SELECTOR, func(h *colly.HTMLElement) {
		h.ForEach(MONTH_SELECTOR, func(i int, monthElement *colly.HTMLElement) {
			monthTitle = strings.TrimSpace(monthElement.Text)
		})
		h.ForEach(DATE_SELECTOR, func(i int, dateElement *colly.HTMLElement) {
			// Go through every cell
			h.ForEach("div#Cell1.cells", func(i int, date *colly.HTMLElement) {
				// count the number of gaps
				if len(date.Attr("style")) > 0 {
					currentDate = endDate + 1
				}
				date.ForEach("div#dashi", func(i int, dashi *colly.HTMLElement) {
					// Only count the gaps in front
					if len(strings.TrimSpace(dashi.Text)) > 0 {
						endDate++
						dontSkipgaps = true
					} else {
						if !dontSkipgaps {
							gaps++
						}
					}
				})
			})
		})
		PrintCalendar(monthTitle, gaps, endDate, currentDate)
	})
	c.Visit(URL)
}

// save the calendar to disk
func PrintCalendar(title string, spaces, end, today int8) {
	// Print the title
	fmt.Println(title)
	// Print separator
	for s := 0; s < 20; s++ {
		fmt.Print("-")
	}
	fmt.Println()
	// Print the name of days, eg sun mon etc
	for _, dayName := range utils.GetDays() {
		fmt.Print(dayName, " ")
	}
	fmt.Print("\b\n")
	for i := int8(0); i < spaces; i++ {
		fmt.Print("   ")
	}

	// Format milayera print gar
	for j := int8(1); j <= end; j++ {
		if (spaces+j)%7 == 0 {
			// Print saturday in red color
			fmt.Printf("\033[1;31m%2d\033[0m", j)
		} else if j == today {
			// Print in blue color
			fmt.Printf("\033[1;34m%2d\033[0m", j)
		} else {
			fmt.Printf("%2d", j)
		}
		fmt.Print(" ")
		if (spaces+j)%7 == 0 {
			fmt.Printf("\b\n")
		}
	}
	fmt.Printf("\n\n")
}

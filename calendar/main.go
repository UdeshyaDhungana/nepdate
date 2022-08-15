package calendar

import (
	"fmt"
	"nepdate/utils"
	"os"

	"github.com/gocolly/colly"
)

const URL = "https://nepalicalendar.rat32.com"
const MONTH_SELECTOR = "html>body>div#page>div#content>div#leftMiddle>div#father>div#cover>div#monthtitle>h1#yren"
const DATE_SELECTOR = "html>body>div#page>div#content>div#leftMiddle>div#father>div#cover>div#main>div#Cells1.cells"

// To store holiday information
type holiday struct {
	name string
	date int8
}

func Calendar() {
	c := colly.NewCollector()

	// On Error
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Could not connect to the network")
		os.Exit(1)
	})
	// On HTML Response
	c.OnHTML(MONTH_SELECTOR, func(h *colly.HTMLElement) {
		// title := strings.TrimSpace(h.Text)
		// Parse gaps, parse end date, parse today, parse holidays
		// PrintCalendar(title, 3, 31, 22)
	})
	// c.Visit(URL)
	var dummyHolidays = []holiday{
		{
			name: "Gaijatra",
			date: 14,
		},
		{
			name: "Dummy Holiday 2",
			date: 30,
		},
	}
	PrintCalendar("Shrawan 2079", 3, 31, 22, dummyHolidays)
}

func PrintCalendar(title string, spaces, end, today int8, holidays []holiday) {
	// Holidays in dates
	var holidayDates []int8

	for _, hDay := range holidays {
		holidayDates = append(holidayDates, hDay.date)
	}

	fmt.Println(title)
	for _, dayName := range utils.GetDays() {
		fmt.Print(dayName, " ")
	}
	fmt.Print("\b\n")
	for i := int8(0); i < spaces; i++ {
		fmt.Print("   ")
	}

	// Format milayera print gar
	for j := int8(1); j <= end; j++ {
		if utils.IsNumberInArray(j, holidayDates) || ((spaces+j)%7 == 0) {
			// Print holiday in red color
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

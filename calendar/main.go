package calendar

import (
	"fmt"
	"nepdate/utils"
	"path/filepath"
	"time"

	// "nepdate/utils"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

const URL = "https://nepalicalendar.rat32.com"
const COMMON_SELECTOR = "html>body>div#page>div#content>div#leftMiddle>div#father>div#cover"
const MONTH_SELECTOR = "div#monthtitle>h1#yren"
const DATE_SELECTOR = "div#main"

func UpdateCache(lastCallPath, calCachePath string) {
	c := colly.NewCollector()
	var monthTitle string
	var gaps int8
	var dontSkipgaps bool
	var endDate int8
	var currentDate int8

	// On Error
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Could not connect to the network")
		os.Exit(1)
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
		// save the cache
		lastCallFile, lcFileError := os.Create(lastCallPath)
		if lcFileError != nil {
			panic("Operation aborted. Please check file permissions")
		}
		fmt.Fprint(lastCallFile, time.Now().Format("2006.01"))

		// write calendar output in
		SaveCache(monthTitle, gaps, endDate, currentDate, calCachePath)
	})
	c.Visit(URL)

}

func Calendar() {
	homeDir := utils.GetHomeDir()
	lastCallPath := filepath.Join(homeDir, ".nepdate", "lastcall_cal")
	calCachePath := filepath.Join(homeDir, ".nepdate", "cal")

	// Read when last called
	lastCalled, err := os.ReadFile(lastCallPath)
	if err != nil {
		UpdateCache(lastCallPath, calCachePath)
	} else {
		// if today's year and month matches
		lc := string(lastCalled)
		tn := time.Now().Format("2006.01")
		if lc != tn {
			// open the file
			UpdateCache(lastCallPath, calCachePath)
		}
	}
	PrintCalendar(lastCallPath, calCachePath)
}

func PrintCalendar(lastCallPath, calCachePath string) {
	cachedDate, err := os.ReadFile(calCachePath)
	if err != nil {
		UpdateCache(lastCallPath, calCachePath)
		PrintCalendar(lastCallPath, calCachePath)
	}
	// Print cache
	fmt.Print(string(cachedDate))
}

func SaveCache(title string, spaces, end, today int8, calCachePath string) {
	calCache, calCacheError := os.Create(calCachePath)
	if calCacheError != nil {
		panic("Operation aborted. Please check file permissions.")
	}
	// Print the title
	fmt.Fprintln(calCache, title)
	// Print separator
	for s := 0; s < 20; s++ {
		fmt.Fprint(calCache, "-")
	}
	fmt.Fprintln(calCache)
	// Print the name of days, eg sun mon etc
	for _, dayName := range utils.GetDays() {
		fmt.Fprint(calCache, dayName, " ")
	}
	fmt.Fprint(calCache, "\b\n")
	for i := int8(0); i < spaces; i++ {
		fmt.Fprint(calCache, "   ")
	}

	// Format milayera print gar
	for j := int8(1); j <= end; j++ {
		if (spaces+j)%7 == 0 {
			// Print saturday in red color
			fmt.Fprintf(calCache, "\033[1;31m%2d\033[0m", j)
		} else if j == today {
			// Print in blue color
			fmt.Fprintf(calCache, "\033[1;34m%2d\033[0m", j)
		} else {
			fmt.Fprintf(calCache, "%2d", j)
		}
		fmt.Fprint(calCache, " ")
		if (spaces+j)%7 == 0 {
			fmt.Fprintf(calCache, "\b\n")
		}
	}
	fmt.Fprintf(calCache, "\n\n")
}

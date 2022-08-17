package converter

import (
	"fmt"
	"nepdate/utils"
	"strings"

	"github.com/gocolly/colly"
)

const CONVERTER_URL = "https://www.bolpatra.gov.np/egp/openDateConverter/convertToEng"
const REV_CONVERTER_URL = "https://www.bolpatra.gov.np/egp/openDateConverter/convertToNep"

// date should be a string of format YYYY.MM.DD
// reverse is true if conversion is from english to nepali date
func GetResults(date string, reverse bool) {
	// check format, if valid, send and receive
	if utils.IsValidDate(date) {
		// variables
		var req string
		var url string

		ymd := utils.GetDMYFromDate(date)
		// go colly request
		if reverse {
			req = fmt.Sprintf("date=%d&month=%d&year=%d", ymd[2], ymd[1], ymd[0])
		} else {
			req = fmt.Sprintf("nepYear=%d&nepMonth=%d&nepDate:%d", ymd[0], ymd[1], ymd[2])
		}

		payload := []byte(req)
		c := colly.NewCollector()
		c.OnRequest(func(r *colly.Request) {
			r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		})
		c.OnHTML("fieldset>table", func(h *colly.HTMLElement) {
			fmt.Print("Converted Date: ")
			fmt.Println(strings.TrimSpace(h.Text))
		})
		if reverse {
			url = REV_CONVERTER_URL
		} else {
			url = CONVERTER_URL
		}
		err := c.PostRaw(url, payload)
		if err != nil {
			panic("There was a network connection error")
		}
	}
}

func Converter(date string) {
	GetResults(date, false)
}

func RevConverter(date string) {
	GetResults(date, true)
}

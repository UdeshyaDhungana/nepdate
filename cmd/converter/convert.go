package converter

import (
	"fmt"
	"io"
	"nepdate/internal"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

const TO_AD_URL = "https://www.bolpatra.gov.np/egp/convertToEng"
const TO_BS_URL = "https://www.bolpatra.gov.np/egp/convertToNep"

func Convert(inDate string, outDateType int) error {
	date, err := ParseDate(inDate)
	if err != nil {
		return fmt.Errorf("Convert: Parsing failed;  %v", err)
	}
	formData := makeFormData(date, outDateType)
	result, err := MakeAPICall(formData, getURL(outDateType))
	if err != nil {
		return fmt.Errorf("ConvertToAD: Api call failed:  %v", err)
	}

	printFormatted(result, outDateType)
	return nil
}

func ParseDate(s string) (Date, error) {
	date, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return Date{}, fmt.Errorf("parseDate: Error parsing %v", err)
	}
	return Date{Year: date.Year(), Month: int(date.Month()), Date: date.Day()}, nil
}

func makeFormData(d Date, dateType int) url.Values {
	date := strconv.Itoa(d.Date)
	month := strconv.Itoa(d.Month)
	year := strconv.Itoa(d.Year)
	switch dateType {
	case ENGLISH_DATE:
		return url.Values{"nepDate": {date}, "nepMonth": {month}, "nepYear": {year}}
	case NEPALI_DATE:
		return url.Values{"date": {date}, "month": {month}, "year": {year}}
	default:
		panic(fmt.Sprintf("makeFormData: Invalid date type: %d", dateType))
	}
}

func getURL(outDateType int) string {
	switch outDateType {
	case ENGLISH_DATE:
		return TO_AD_URL
	case NEPALI_DATE:
		return TO_BS_URL
	default:
		panic(fmt.Sprintf("getURL: Invalid outDateType %d", outDateType))
	}
}

const CONTENT_TYPE = "application/x-www-form-urlencoded"
const CSS_SELECTOR = "fieldset>table>tbody>tr>td"

// might need to split this into requester, parser, and selector
// the web interface provides same html template for both requests, so i'm using a single function for nows
func MakeAPICall(formData url.Values, requestUrl string) (string, error) {
	resp, err := http.PostForm(requestUrl, formData)
	if err != nil {
		return "", fmt.Errorf("MakeAPICall: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("MakeAPICall: %v", resp.Status)
	}

	// read body as bytes
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("MakeAPICall: Error reading response body: %v", err)
	}

	// convert to string, and parse html
	bodyString := string(bodyBytes)
	doc, err := html.Parse(strings.NewReader(bodyString))
	if err != nil {
		return "", fmt.Errorf("MakeAPICall: Error parsing response as HTML: %v", err)
	}

	// prepare cascadia css selector
	sel, err := cascadia.Parse(CSS_SELECTOR)
	if err != nil {
		return "", fmt.Errorf("MakeAPICall: Error parsing css selector %s", CSS_SELECTOR)
	}

	// scrape
	dateTd := cascadia.Query(doc, sel)
	if dateTd == nil || dateTd.FirstChild.Type != html.TextNode {
		return "", fmt.Errorf("MakeAPICall: Error using CSS Selector: %v", err)
	}

	return strings.TrimSpace(dateTd.FirstChild.Data), nil
}

// Expected input is as is for result
// English: Friday, December 15, 2000
// Nepali: Friday, 2057/08/30
// So, convert nepali date to month name
func printFormatted(result string, outDateType int) error {
	switch outDateType {
	case ENGLISH_DATE:
		fmt.Println(result)
	case NEPALI_DATE:
		s1 := strings.Split(result, ", ")
		s2 := strings.Split(s1[1], "/")

		weekDay := s1[0]
		monthFromResult := s2[1]
		monthIndex, err := strconv.Atoi(monthFromResult)
		if err != nil {
			return fmt.Errorf("printFormatted: Failed to parse response date")
		}
		monthIndex -= 1
		monthName := internal.GetMonthName(monthIndex)
		year := s2[0]
		date := s2[2]

		fmt.Printf("%s, %s %s, %s\n", weekDay, monthName, date, year)
	default:
		panic(fmt.Sprintf("printFormatted: Invalid outDateType: %d", outDateType))
	}
	return nil
}

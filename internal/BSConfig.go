package internal

import (
	"fmt"
	"nepdate/internal/utils"
	"time"
)

// BS dates dataset/config; call it whatever you like
var bsConfig = map[int]YearInfo{
	2079: {
		Year:         2079,
		DaysInMonths: [_MONTHS]int{31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
		StartDate:    1649894400,
	},
	2080: {
		Year:         2080,
		DaysInMonths: [_MONTHS]int{31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
		StartDate:    1681430400,
	},
}

type YearInfo struct {
	// Nepali year
	Year int
	// Array containing the number of days in each month
	DaysInMonths [_MONTHS]int
	// Unix timestamp for Baisakh 1
	StartDate int64
}

const _MONTHS = 12

var _MONTH_NAMES = map[int]string{
	0:  "Baisakh",
	1:  "Jestha",
	2:  "Ashar",
	3:  "Shrawan",
	4:  "Bhadra",
	5:  "Ashwin",
	6:  "Kartik",
	7:  "Mangsir",
	8:  "Poush",
	9:  "Magh",
	10: "Falgun",
	11: "Chaitra",
}

var NO_CONFIG_ERROR = "We don't have config for %d B.S. Please report to developers"

func GetMonthName(idx int) string {
	month, ok := _MONTH_NAMES[idx]
	if !ok {
		panic("Invalid month index")
	}
	return month
}

func GetCurrentYearInfo() YearInfo {
	nowADYear := time.Now().Year()
	plus56, ok := bsConfig[nowADYear+56]
	if !ok {
		panic(fmt.Sprintf(NO_CONFIG_ERROR, nowADYear))
	}
	plus57, ok := bsConfig[nowADYear+57]
	if !ok {
		panic(fmt.Sprintf(NO_CONFIG_ERROR, nowADYear))
	}
	if time.Now().Unix()-plus57.StartDate > 0 {
		return plus57
	} else {
		return plus56
	}
}

// Returns the nepali date
// if month is negative, otherwise return today's nepali date
func GetNepaliDate(year int, month int, date int) (NepaliDate, error) {
	if month > 11 {
		return NepaliDate{}, fmt.Errorf("GetNepaliDate: month out of index")
	}
	yearInfo, ok := bsConfig[year]
	if !ok {
		return NepaliDate{}, fmt.Errorf(NO_CONFIG_ERROR, year)
	}
	newYearTS, daysInMonths := yearInfo.StartDate, yearInfo.DaysInMonths
	newYearTime := time.Unix(newYearTS, 0)

	// Calculate date for today
	if month < 0 {
		daysSinceNewYear := utils.DaysSince((newYearTime))
		todayTS := newYearTime.AddDate(0, 0, daysSinceNewYear).Unix()
		date := utils.DaysSinceInclusive(newYearTime)
		for i, dayInMonth := range daysInMonths {
			date -= dayInMonth
			if date < 0 {
				return NepaliDate{Year: year, Month: i, Date: date + dayInMonth, Timestamp: todayTS}, nil
			}
		}
	}
	// Calculate for the specified month and date
	if daysInMonths[month] < date {
		return NepaliDate{}, fmt.Errorf("the month does not has the date")
	}
	daysSinceNewYearInclusive := date
	for i := 0; i < month; i++ {
		daysSinceNewYearInclusive += daysInMonths[i]
	}
	thatDaysTS := newYearTime.AddDate(0, 0, daysSinceNewYearInclusive-1).Unix()
	return NepaliDate{Year: year, Month: month, Date: date, Timestamp: thatDaysTS}, nil
}

func GetCurrentDate() (NepaliDate, error) {
	currentYearInfo := GetCurrentYearInfo()
	currentDate, error := GetNepaliDate(currentYearInfo.Year, -1, -1)
	if error != nil {
		return NepaliDate{}, error
	}
	return currentDate, nil
}

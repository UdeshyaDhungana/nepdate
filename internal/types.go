package internal

import (
	"fmt"
	"math"
	"time"
)

type BSMonth int

const (
	Baisakh BSMonth = iota
	Jestha
	Ashar
	Shrawan
	Bhadra
	Asoj
	Kartik
	Mangsir
	Poush
	Magh
	Falgun
	Chaitra
)

func (m BSMonth) String() string {
	switch m {
	case Baisakh:
		return "Baisakh"
	case Jestha:
		return "Jestha"
	case Ashar:
		return "Ashar"
	case Shrawan:
		return "Shrawan"
	case Bhadra:
		return "Bhadra"
	case Asoj:
		return "Asoj"
	case Kartik:
		return "Kartik"
	case Mangsir:
		return "Mangsir"
	case Poush:
		return "Poush"
	case Magh:
		return "Magh"
	case Falgun:
		return "Falgun"
	case Chaitra:
		return "Chaitra"
	}
	panic(fmt.Sprintf("Invalid BSMonth %d", m))
}

type BSDate struct {
	Year  int
	Month BSMonth
	Day   int
}

func (d BSDate) Weekday() time.Weekday {
	return MustConvertBsDateToAdTime(d).Weekday()
}

func (d BSDate) GetSakrati() BSDate {
	return BSDate{
		Year:  d.Year,
		Month: d.Month,
		Day:   1,
	}
}

func (d BSDate) String() string {
	return fmt.Sprintf("%s %s %02d %d", d.Weekday(), d.Month, d.Day, d.Year)
}

// Converts *time.Time to Bikram Samvat Date
// If the AD equivalent time doesn't fall within our config's limits, it panics

func MustConvertAdTimeToBsDate(time time.Time) BSDate {
	timeUnixTs := time.Unix()
	bsConfig := GetConfig()
	minPositiveDiffBetweenNowAndNewYear := int64(math.MaxInt64)
	var bsYearConfig BSYearConfig
	var bsYear int

	for year, yearConfig := range bsConfig.Years {
		yearStartTs := yearConfig.StartDateUnixTs
		difference := timeUnixTs - yearStartTs

		if difference >= 0 && difference < minPositiveDiffBetweenNowAndNewYear {
			minPositiveDiffBetweenNowAndNewYear = difference
			bsYearConfig = yearConfig
			bsYear = year
		}
	}

	daysSinceNayaBarsa := 1 + (timeUnixTs-bsYearConfig.StartDateUnixTs)/86400
	totalDaysInYear := 0

	for _, daysInMonth := range bsYearConfig.DaysInMonths {
		totalDaysInYear += daysInMonth
	}

	if daysSinceNayaBarsa > int64(totalDaysInYear) {
		panic("Date falls outside of configured bounds.\nRaise a PR here: https://github.com/UdeshyaDhungana/nepdate/blob/main/internal/years.json")
	}

	monthsSinceBSNewYear := 0
	daysSinceSakrati := daysSinceNayaBarsa

	for _, daysInMonth := range bsYearConfig.DaysInMonths {
		if daysSinceSakrati-int64(daysInMonth) <= 0 {
			break
		}
		daysSinceSakrati -= int64(daysInMonth)
		monthsSinceBSNewYear += 1
	}

	return BSDate{
		Year:  bsYear,
		Month: BSMonth(monthsSinceBSNewYear),
		Day:   int(daysSinceSakrati),
	}
}

func MustConvertBsDateToAdTime(bsDate BSDate) time.Time {
	config := GetConfig()

	bsYearConfig, ok := config.Years[bsDate.Year]
	if !ok {
		panic("Date falls outside of configured bounds.\nRaise a PR here: https://github.com/UdeshyaDhungana/nepdate/blob/main/internal/years.json")
	}

	daysSinceNayaBarsa := 0

	for _, daysInMonth := range bsYearConfig.DaysInMonths[:bsDate.Month] {
		daysSinceNayaBarsa += daysInMonth
	}

	daysSinceNayaBarsa += bsDate.Day - 1
	nowTs := bsYearConfig.StartDateUnixTs + int64(86400*daysSinceNayaBarsa)
	return time.Unix(nowTs, 0)
}

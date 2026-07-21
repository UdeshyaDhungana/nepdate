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

	// This algorithm suffers from one important edge case
	// If the config isn't updated in the future
	// current year will be set to the max possible year whose config is set
	// If you come across such a mistake, raise a PR adding the config
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
	monthsSinceBSNewYear := 0

	daysSinceSakrati := daysSinceNayaBarsa
	for _, daysInMonth := range bsYearConfig.DaysInMonths {
		if daysSinceSakrati-int64(daysInMonth) >= 0 {
			daysSinceSakrati -= int64(daysInMonth)
			monthsSinceBSNewYear += 1
		}
	}

	return BSDate{
		Year:  bsYear,
		Month: BSMonth(monthsSinceBSNewYear),
		Day:   int(daysSinceSakrati),
	}
}

func MustConvertBsDateToAdTime(bsDate BSDate) time.Time {
	config := GetConfig()
	bsYearConfig := config.Years[bsDate.Year]
	daysSinceNayaBarsa := 0

	for _, daysInMonth := range bsYearConfig.DaysInMonths[:bsDate.Month] {
		daysSinceNayaBarsa += daysInMonth
	}

	daysSinceNayaBarsa += bsDate.Day - 1
	nowTs := bsYearConfig.StartDateUnixTs + int64(86400*daysSinceNayaBarsa)
	return time.Unix(nowTs, 0)
}

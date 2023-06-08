package internal

import (
	"fmt"
	"time"
)

type NepaliDate struct {
	// Nepali year
	Year int
	// Index of month (0 - 11)
	Month int
	// Date of current month
	Date int
	// Unix timestamp for 00:00:00
	Timestamp int64
}

func (n NepaliDate) Print() {
	month := GetMonthName(n.Month)
	fmt.Printf("%02d %s %d, %s\n", n.Date, month, n.Year, time.Unix(n.Timestamp, 0).Weekday())
}

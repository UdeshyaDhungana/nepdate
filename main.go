package main

import (
	"flag"
	"nepdate/calendar"
	"nepdate/converter"
	"nepdate/date"
	"os"
)

func main() {
	args := os.Args[1:]
	var calFlag = flag.Bool("cal", false, "Use -cal to invoke a calendar.")
	var dateFlag = flag.Bool("date", false, "Use -date or no arguments for today's date.")
	var convFlag = flag.String("conv", "", "Use -conv=YYYY.MM.DD to convert to English.")
	var revConvFlag = flag.String("rconv", "", "Use -rconv=YYYY.MM.DD to convert to Nepalese.")
	flag.Parse()

	if len(args) == 0 {
		date.Date()
		// fmt.Println("Executing date...")
	} else {
		// If dateFlag is present, show date
		if *dateFlag {
			date.Date()
			// fmt.Println("Executing date...")
		}
		// If calFlag is present, show calendar
		if *calFlag {
			// call calendar command
			calendar.Calendar()
		}
		// If convflag is present, show converted date
		if len(*convFlag) > 0 {
			converter.Converter(*convFlag)
		}
		// If revConvFlag is present, show it here
		if len(*revConvFlag) > 0 {
			converter.RevConverter(*revConvFlag)
		}
	}
}

package main

import (
	"flag"
	"nepdate/calendar"
	"nepdate/date"
	"os"
)

func main() {
	args := os.Args[1:]
	var calFlag = flag.Bool("cal", false, "Use -cal to invoke a calendar.")
	var dateFlag = flag.Bool("date", false, "Use -date or no arguments for today's date.")
	flag.Parse()

	if len(args) == 0 {
		date.Date()
		// fmt.Println("Executing date...")
	} else {
		if *calFlag {
			// call calendar command
			calendar.Calendar()
		} else if *dateFlag {
			date.Date()
			// fmt.Println("Executing date...")
		}
	}
}

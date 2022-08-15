package main

import (
	"flag"
	"fmt"
	"nepdate/calendar"
	"os"
)

func main() {
	args := os.Args[1:]
	var calFlag = flag.Bool("cal", false, "Use -cal or --cal to invoke a calendar.")
	var dateFlag = flag.Bool("date", false, "Use no arguments or -date for today's date.")
	flag.Parse()

	if len(args) == 0 {
		// date.Date()
		fmt.Println("Executing date...")
	} else {
		if *calFlag {
			// call calendar command
			calendar.Calendar()
		} else if *dateFlag {
			// date.Date()
			fmt.Println("Executing date...")
		}
	}
}

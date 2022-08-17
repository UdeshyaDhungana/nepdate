package main

import (
	"flag"
	"fmt"
	"nepdate/calendar"
	"nepdate/converter"
	"nepdate/date"
	"nepdate/utils"
	"os"
)

func main() {
	args := os.Args[1:]
	helpString := "Use -help for instructions."
	var calFlag = flag.Bool("cal", false, helpString)
	var dateFlag = flag.Bool("date", false, helpString)
	var convFlag = flag.String("conv", "", helpString)
	var revConvFlag = flag.String("rconv", "", helpString)
	var helpFlag = flag.Bool("help", false, helpString)
	flag.Parse()

	if len(args) == 0 {
		date.Date()
		// fmt.Println("Executing date...")
	} else {
		if *helpFlag {
			fmt.Printf("Nepdate is a tool for working with Nepalese dates.\n\n")
			fmt.Printf("Usage: nepdate [options]\n\n")
			fmt.Printf("The options are:\n\n")
			helpOptions := utils.GetHelp()
			for _, helpOption := range helpOptions {
				fmt.Printf("\t%s\n", helpOption.Option)
				fmt.Printf("\t%s\n\n", helpOption.HelpText)
			}
		} else
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

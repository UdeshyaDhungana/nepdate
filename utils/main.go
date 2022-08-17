package utils

import (
	"os"
	"strconv"
	"time"
)

// Struct for working with helpMessage
type helpMessage struct {
	Option   string
	HelpText string
}

func GetDays() []string {
	return []string{"Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"}
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Error accessing user's home directory")
	}
	return homeDir
}

func IsNumberInArray(a int8, list []int8) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IsValidDate(date string) bool {
	_, err := time.Parse("2006.01.02", date)
	if err != nil {
		panic("Aborted: Date is invalid.")
	}
	return (err == nil)
}

// extract [YYY, MM, DD] from YYYY.MM.DD
func GetDMYFromDate(date string) []int {
	year, yerr := strconv.Atoi(date[0:4])
	month, merr := strconv.Atoi(date[5:7])
	day, derr := strconv.Atoi(date[8:10])
	if yerr != nil || merr != nil || derr != nil {
		panic("Aborted: Error while parsing date.")
	}
	return []int{
		year,
		month,
		day,
	}
}

// get help options
func GetHelp() []helpMessage {
	helpMessages := []helpMessage{
		{
			Option:   "-date",
			HelpText: "Displays today's Nepalese date",
		},
		{
			Option:   "-cal",
			HelpText: "Display this month's Nepalese calendar",
		},
		{
			Option:   "-conv=YYYY.MM.DD",
			HelpText: "Convert and display the specified Nepalese date to English",
		},
		{
			Option:   "-rconv=YYYY.MM.DD",
			HelpText: "Convert and display the specified English date to Nepalese",
		},
		{
			Option:   "-help",
			HelpText: "Display this help meessage",
		},
	}
	return helpMessages
}

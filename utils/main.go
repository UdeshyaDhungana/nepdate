package utils

import (
	"os"
)

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
